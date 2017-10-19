package zonedb

import "testing"

type runePointsTest struct {
	C      rune
	Points []rune
	Result bool
}

var asciiRunePointsData = []runePointsTest{
	{C: 0, Points: asciiCodePoints, Result: false},
	{C: ' ', Points: asciiCodePoints, Result: false},
	{C: '-', Points: asciiCodePoints, Result: true},
	{C: '/', Points: asciiCodePoints, Result: false},
	{C: '0', Points: asciiCodePoints, Result: true},
	{C: '9', Points: asciiCodePoints, Result: true},
	{C: ':', Points: asciiCodePoints, Result: false},
	{C: '`', Points: asciiCodePoints, Result: false},
	{C: 'a', Points: asciiCodePoints, Result: true},
	{C: 'b', Points: asciiCodePoints, Result: true},
	{C: 'z', Points: asciiCodePoints, Result: true},
	{C: '{', Points: asciiCodePoints, Result: false},
	{C: '~', Points: asciiCodePoints, Result: false},
	{C: 'é', Points: asciiCodePoints, Result: false},
	{C: 'ø', Points: asciiCodePoints, Result: false},
}

var unicodeCodePoints = []rune{'-', '-', '0', '9', 'а', 'ъ', 'ь', 'ь', 'ю', 'я'}

var unicodeRunePointsData = []runePointsTest{
	{C: 0, Points: unicodeCodePoints, Result: false},
	{C: 'з', Points: unicodeCodePoints, Result: true},
	{C: 'д', Points: unicodeCodePoints, Result: true},
	{C: 'р', Points: unicodeCodePoints, Result: true},
	{C: 'а', Points: unicodeCodePoints, Result: true},
	{C: 'в', Points: unicodeCodePoints, Result: true},
	{C: 'е', Points: unicodeCodePoints, Result: true},
	{C: 'й', Points: unicodeCodePoints, Result: true},
	{C: 'в', Points: unicodeCodePoints, Result: true},
	{C: 'с', Points: unicodeCodePoints, Result: true},
	{C: 'ё', Points: unicodeCodePoints, Result: true},
	{C: '{', Points: unicodeCodePoints, Result: false},
	{C: '~', Points: unicodeCodePoints, Result: false},
}

var manyUnicodeCodePoints = []rune{'-', '-', '0', '9', 'À', 'Á', 'Â', 'Ã', 'Ä', 'Å', 'Æ', 'Ç', 'È', 'É', 'Ê', 'Ë', 'Ì', 'Í', 'Î', 'Ï', 'Ð', 'Ñ', 'Ò', 'Ó', 'Ô', 'Õ', 'Ö', 'Ø', 'Ù', 'Ú', 'Û', 'Ü', 'Ý', 'Þ', 'Ā', 'Ă', 'Ą', 'Ć', 'Ĉ', 'Ċ', 'Č', 'Ď', 'Đ', 'Ē', 'Ĕ', 'Ė', 'Ę', 'Ě', 'Ĝ', 'Ğ', 'Ġ', 'Ģ', 'Ĥ', 'Ħ', 'Ĩ', 'Ī', 'Ĭ', 'Į', 'İ', 'Ĳ', 'Ĵ', 'Ķ', 'Ĺ', 'Ļ', 'Ľ', 'Ŀ', 'Ł', 'Ń', 'Ņ', 'Ň', 'Ŋ', 'Ō', 'Ŏ', 'Ő', 'Œ', 'Ŕ', 'Ŗ', 'Ř', 'Ś', 'Ŝ', 'Ş', 'Š', 'Ţ', 'Ť', 'Ŧ', 'Ũ', 'Ū', 'Ŭ', 'Ů', 'Ű', 'Ų', 'Ŵ', 'Ŷ', 'Ÿ', 'Ź', 'Ż', 'Ž', 'Ɓ', 'Ƃ', 'Ƅ', 'Ɔ', 'Ƈ', 'Ɖ', 'Ɗ', 'Ƌ', 'Ǝ', 'Ə', 'Ɛ', 'Ƒ', 'Ɠ', 'Ɣ', 'Ɩ', 'Ɨ', 'Ƙ', 'Ɯ', 'Ɲ', 'Ɵ', 'Ơ', 'Ƣ', 'Ƥ', 'Ʀ', 'Ƨ', 'Ʃ', 'Ƭ', 'Ʈ', 'Ư', 'Ʊ', 'Ʋ', 'Ƴ', 'Ƶ', 'Ʒ', 'Ƹ', 'Ƽ', 'Ǆ', 'Ǉ', 'Ǌ', 'Ǎ', 'Ǐ', 'Ǒ', 'Ǔ', 'Ǖ', 'Ǘ', 'Ǚ', 'Ǜ', 'Ǟ', 'Ǡ', 'Ǣ', 'Ǥ', 'Ǧ', 'Ǩ', 'Ǫ', 'Ǭ', 'Ǯ', 'Ǳ', 'Ǵ', 'Ƕ', 'Ƿ', 'Ǹ', 'Ǻ', 'Ǽ', 'Ǿ', 'Ȁ', 'Ȃ', 'Ȅ', 'Ȇ', 'Ȉ', 'Ȋ', 'Ȍ', 'Ȏ', 'Ȑ', 'Ȓ', 'Ȕ', 'Ȗ', 'Ș', 'Ȟ', 'Ƞ', 'Ȣ', 'Ȥ', 'Ȧ', 'Ȩ', 'Ȫ', 'Ȭ', 'Ȯ', 'Ȱ'}

var manyUnicodeRunePointsData = []runePointsTest{
	{C: 0, Points: manyUnicodeCodePoints, Result: false},
	{C: 'Þ', Points: manyUnicodeCodePoints, Result: true},
	{C: 'Ê', Points: manyUnicodeCodePoints, Result: true},
	{C: 'Ñ', Points: manyUnicodeCodePoints, Result: true},
	{C: 'Ȣ', Points: manyUnicodeCodePoints, Result: true},
	{C: 'Ȫ', Points: manyUnicodeCodePoints, Result: true},
	{C: 'Ȝ', Points: manyUnicodeCodePoints, Result: false},
	{C: 'й', Points: manyUnicodeCodePoints, Result: false},
	{C: 'в', Points: manyUnicodeCodePoints, Result: false},
	{C: 'с', Points: manyUnicodeCodePoints, Result: false},
	{C: 'ё', Points: manyUnicodeCodePoints, Result: false},
	{C: '{', Points: manyUnicodeCodePoints, Result: false},
	{C: '~', Points: manyUnicodeCodePoints, Result: false},
}

func TestRuneInCodePoints(t *testing.T) {
	for _, d := range asciiRunePointsData {
		r := runeInCodePoints(d.C, d.Points)
		if r != d.Result {
			t.Errorf(`Expected runeInCodePoints(%q, %q) == %t, got %t`, d.C, d.Points, d.Result, r)
		}
	}
}
