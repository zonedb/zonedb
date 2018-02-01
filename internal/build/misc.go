package build

// IndexOfStrings finds a rank-sorted slice of domain names (needle)
// within a larger slice (haystack).
func IndexOfStrings(haystack []string, needle []string) int {
outer:
	for i := range haystack {
		for j := range needle {
			if i+j >= len(haystack) || needle[j] != haystack[i+j] {
				continue outer
			}
		}
		return i
	}
	return -1
}

// IndexOrAppendStrings finds or appends a slice of rank-sorted domain names (needle)
// Returns 0,0 for a zero-length needle.
func IndexOrAppendStrings(haystack *[]string, needle []string) (int, int) {
	if len(needle) == 0 {
		return 0, 0
	}
	idx := IndexOfStrings(*haystack, needle)
	if idx < 0 {
		idx = len(*haystack)
		*haystack = append(*haystack, needle...)
	}
	return idx, idx + len(needle)
}
