package slidingwindow

import "slices"

func StrStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	target := []byte(needle)

	// We need a window of size needle
	window := make([]byte, len(needle))

	// Initialize the window over haystack
	for i := 0; i < len(window); i++ {
		window[i] = haystack[i]
	}

	// Do a prelim check for the needle
	if slices.Equal(window, target) {
		return 0
	}

	// Slide the window over the haystack until target is found
	for i := 1; i <= len(haystack)-len(window); i++ {
		sample := haystack[i : i+len(window)]
		if slices.Equal([]byte(sample), target) {
			return i
		}
	}

	return -1
}
