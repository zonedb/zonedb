package zonedb

var asciiCodePoints = []rune("--09az")

// StringInCodePoints determines if a string falls within the range(s) specified in points.
// Points must be an even-numbered set of sorted Unicode code points (low, high).
// Note: This is an experimental API and may change.
func StringInCodePoints(s string, points []rune) bool {
	min := rune('\U0010FFFF')
	max := rune(0)

	for _, c := range s {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}

	if !runeInCodePoints(max, points) {
		return false
	}
	if !runeInCodePoints(min, points) {
		return false
	}

	for _, c := range s {
		if !runeInCodePoints(c, points) {
			return false
		}
	}
	return true
}

func runeInCodePoints(c rune, points []rune) bool {
	return linearRuneInCodePoints(c, points)
}

func linearRuneInCodePoints(c rune, points []rune) bool {
	end := len(points) - 1
	if c > points[end] || c < points[0] {
		return false
	}
	for i := 0; i < end; i += 2 {
		if c >= points[i] && c <= points[i+1] {
			return true
		}
	}
	return false
}

func binaryRuneInCodePoints(c rune, points []rune) bool {
	var min, midMin, max int

	min = 0
	max = len(points) - 1

	if c > points[max] || c < points[min] {
		return false
	}

	for min < max-1 {
		midMin = min + (((max - min) >> 1) << 1)
		if c < points[midMin] {
			max = midMin - 1
		} else {
			min = midMin
		}
	}
	return c >= points[min] && c <= points[max]
}
