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

var unicodeCodePoints = []rune("--09аъььюя")

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
	{C: 'ё', Points: unicodeCodePoints, Result: false},
	{C: '{', Points: unicodeCodePoints, Result: false},
	{C: '~', Points: unicodeCodePoints, Result: false},
}

var manyUnicodeCodePoints = []rune("--09ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖØÙÚÛÜÝÞĀĂĄĆĈĊČĎĐĒĔĖĘĚĜĞĠĢĤĦĨĪĬĮİĲĴĶĹĻĽĿŁŃŅŇŊŌŎŐŒŔŖŘŚŜŞŠŢŤŦŨŪŬŮŰŲŴŶŸŹŻŽƁƂƄƆƇƉƊƋƎƏƐƑƓƔƖƗƘƜƝƟƠƢƤƦƧƩƬƮƯƱƲƳƵƷƸƼǄǇǊǍǏǑǓǕǗǙǛǞǠǢǤǦǨǪǬǮǱǴǶǷǸǺǼǾȀȂȄȆȈȊȌȎȐȒȔȖȘȞȠȢȤȦȨȪȬȮȰ")

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

var runePointsData = append(append(asciiRunePointsData, unicodeRunePointsData...), manyUnicodeRunePointsData...)

func TestRuneInCodePoints(t *testing.T) {
	for _, d := range runePointsData {
		r := runeInCodePoints(d.C, d.Points)
		if r != d.Result {
			t.Errorf(`Expected runeInCodePoints(%q, %q) == %t, got %t`, d.C, d.Points, d.Result, r)
		}
	}
}

func BenchmarkManyUnicodeLinearRuneInCodePoints(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, d := range manyUnicodeRunePointsData {
			linearRuneInCodePoints(d.C, d.Points)
		}
	}
}

func BenchmarkManyUnicodeBinaryRuneInCodePoints(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, d := range manyUnicodeRunePointsData {
			binaryRuneInCodePoints(d.C, d.Points)
		}
	}
}

func BenchmarkUnicodeLinearRuneInCodePoints(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, d := range unicodeRunePointsData {
			linearRuneInCodePoints(d.C, d.Points)
		}
	}
}

func BenchmarkUnicodeBinaryRuneInCodePoints(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, d := range unicodeRunePointsData {
			binaryRuneInCodePoints(d.C, d.Points)
		}
	}
}

func BenchmarkASCIILinearRuneInCodePoints(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, d := range asciiRunePointsData {
			linearRuneInCodePoints(d.C, d.Points)
		}
	}
}

func BenchmarkASCIIBinaryRuneInCodePoints(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, d := range asciiRunePointsData {
			binaryRuneInCodePoints(d.C, d.Points)
		}
	}
}
