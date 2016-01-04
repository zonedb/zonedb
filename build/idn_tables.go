package build

import (
	"bufio"
	"bytes"
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
			res, err := Fetch(url)
			if err != nil {
				LogWarning(err)
				continue
			}
			defer res.Body.Close()

			table, err := ProcessIDNTable(res.Body, strings.HasSuffix(url, ".html"))
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
	return nil
}

func ProcessIDNTable(data io.Reader, isHTML bool) (*CodeTable, error) {
	unicodePrefix := []byte("U+")
	hexPrefix := []byte("0x")
	dot := byte('.')
	space := byte(' ')
	semi := byte(';')

	if isHTML {
		doc, err := goquery.NewDocumentFromReader(data)
		if err != nil {
			LogWarning(err)
			return nil, err
		}
		data = strings.NewReader(doc.FindMatcher(preMatcher).Text())
	}

	scanner := bufio.NewScanner(data)
	table := &CodeTable{}
	for scanner.Scan() {
		line := scanner.Bytes()

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
			if isHexByte(line[firstHexOffset+firstHexLen]) {
				firstHexLen = 5
			}

			codePoint, err := strconv.ParseInt(string(line[firstHexOffset:firstHexOffset+firstHexLen]), 16, 32)
			if err != nil {
				LogWarning(err)
				continue
			}
			charRune := rune(codePoint)

			append := true
			// Handle U+####..U+#### ranges seen in some RFC 4290 tables
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
