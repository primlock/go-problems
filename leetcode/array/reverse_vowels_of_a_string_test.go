package array

import "testing"

func TestReverseVowels(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected string
	}{
		{name: "1", s: "IceCreAm", expected: "AceCreIm"},
		{name: "2", s: "leetcode", expected: "leotcede"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseVowels(tt.s)
			want := tt.expected

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}
