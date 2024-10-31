package slidingwindow

import "testing"

func TestMaxVowels(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{name: "1", s: "abciiidef", k: 3, expected: 3},
		{name: "2", s: "aeiou", k: 2, expected: 2},
		{name: "3", s: "leetcode", k: 3, expected: 2},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxVowels(tt.s, tt.k)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
