package slidingwindow

import "testing"

func TestStrStr(t *testing.T) {
	testCases := []struct {
		name     string
		needle   string
		haystack string
		expected int
	}{
		{name: "1", needle: "sad", haystack: "sadbutsad", expected: 0},
		{name: "2", needle: "leeto", haystack: "leetcode", expected: -1},
		{name: "3", needle: "issi", haystack: "mississippi", expected: 1},
		{name: "4", needle: "ll", haystack: "hello", expected: 2},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := StrStr(tt.haystack, tt.needle)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
