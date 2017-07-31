package build

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/andybalholm/cascadia"
)

var preMatcher = cascadia.MustCompile("body > pre")

func FetchIDNTables(zones map[string]*Zone) error {
	for _, z := range zones {
		if len(z.IDNTableURLs) < 1 {
			continue
		}
		for lang, url := range z.IDNTableURLs {
			table, err := cachedFetchProcessIDNTable(url)
			if err != nil {
				LogWarning(err)
				continue
			}

			if len(z.IDNTables) == 0 {
				z.IDNTables = make(map[string]CodeTable)
			}
			z.IDNTables[lang] = *table
			z.CodePoints.AppendTable(table)
		}
		z.CodePoints.Compress()
	}
	Trace("@{.}FetchIDNTables: updated %d zones from %d distinct codetable URLs\n", len(zones), len(cacheURLtoCodeTable))
	return nil
}

// If the value stored is nil, then we failed to parse it once, don't re-fetch on errors
var cacheURLtoCodeTable map[string]*CodeTable

var ErrAlreadyFailedToParseOnce = errors.New("already failed to parse this URL once") // should this be a type with the URL in it?

func init() {
	cacheURLtoCodeTable = make(map[string]*CodeTable, 1000)
}

func cachedFetchProcessIDNTable(url string) (*CodeTable, error) {
	if ct, ok := cacheURLtoCodeTable[url]; ok {
		if ct == nil {
			return nil, ErrAlreadyFailedToParseOnce
		}
		return ct.Dup(), nil
	}

	res, err := Fetch(url)
	if err != nil {
		cacheURLtoCodeTable[url] = nil
		return nil, err
	}
	defer res.Body.Close()
	table, err := ProcessIDNTable(res.Body, strings.HasSuffix(url, ".html"), url)
	if err == nil {
		cacheURLtoCodeTable[url] = table
		return table.Dup(), nil
	}
	cacheURLtoCodeTable[url] = nil
	return table, err
}

func ProcessIDNTable(data io.Reader, isHTML bool, label string) (*CodeTable, error) {
	unicodePrefix := []byte("U+")
	hexPrefix := []byte("0x")
	dot := byte('.')
	space := byte(' ')
	semi := byte(';')

	if isHTML {
		doc, err := goquery.NewDocumentFromReader(data)
		if err != nil {
			LogWarningFor(err, label)
			return nil, err
		}
		data = strings.NewReader(doc.FindMatcher(preMatcher).Text())
	}

	scanner := bufio.NewScanner(data)
	table := &CodeTable{}
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		lineNum += 1

		lineLen := len(line)
		if lineLen < 6 {
			continue
		}

		offset := 0
		for line[offset] == space && offset < lineLen-1 {
			offset = offset + 1
		}

		matchedLine := false

		prefixLen := 2
		firstTwo := line[offset : offset+prefixLen]

		// Most tables use U+0000, but some include 0x0000.
		if bytes.Equal(firstTwo, unicodePrefix) || bytes.Equal(firstTwo, hexPrefix) {
			matchedLine = true

			// Some have no prefix, but are four hex chars followed by ;
		} else if line[offset+4] == semi && isHexBytes(line[offset:offset+4]) {
			matchedLine = true
			prefixLen = 0
		}

		if matchedLine {
			firstHexOffset := offset + prefixLen

			// Most are four chars long, but some are five.
			firstHexLen := 4
			if lackSpace(line, firstHexOffset+firstHexLen+1, label, lineNum) {
				continue
			}
			if isHexByte(line[firstHexOffset+firstHexLen]) {
				firstHexLen = 5
			}

			codePoint, err := strconv.ParseInt(string(line[firstHexOffset:firstHexOffset+firstHexLen]), 16, 32)
			if err != nil {
				LogWarningForAt(err, label, lineNum)
				continue
			}
			charRune := rune(codePoint)

			append := true
			// Handle U+####..U+#### ranges seen in some RFC 4290 tables
			if lackSpace(line, firstHexOffset+firstHexLen+2, label, lineNum) {
				continue
			}
			if line[firstHexOffset+firstHexLen+1] == dot {
				secondOffset := firstHexOffset + firstHexLen + 2
				matchedSecond := false
				if prefixLen > 0 {
					secondTwo := line[secondOffset : secondOffset+prefixLen]
					if bytes.Equal(secondTwo, unicodePrefix) || bytes.Equal(secondTwo, hexPrefix) {
						matchedSecond = true
					}
				} else if isHexBytes(line[secondOffset : secondOffset+4]) {
					matchedSecond = true
				}

				if matchedSecond {
					secondHexOffset := secondOffset + prefixLen
					secondHexLen := 4
					if isHexByte(line[secondHexOffset+secondHexLen]) {
						secondHexLen = 5
					}
					endCodePoint, err := strconv.ParseInt(string(line[secondHexOffset:secondHexOffset+secondHexLen]), 16, 32)
					if err == nil {
						append = false
						endCharRune := rune(endCodePoint)
						table.AppendRange(charRune, endCharRune)
					}
				}
			}
			if append {
				table.AppendRune(charRune)
			}
		}
	}
	table.Compress()

	return table, nil
}

func isHexBytes(s []byte) bool {
	for _, c := range s {
		if !isHexByte(c) {
			return false
		}
	}
	return true
}

func isHexByte(c byte) bool {
	return (c >= 48 && c <= 57) || (c >= 65 && c <= 70) || (c >= 97 && c <= 102)
}

func lackSpace(line []byte, need int, label string, lineNum int) bool {
	if len(line) >= need {
		return false
	}
	LogWarningForAt(fmt.Errorf("line too short, have %d need %d", len(line), need), label, lineNum)
	return true
}
