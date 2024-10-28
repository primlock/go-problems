package array

import "testing"

func TestReverseWords(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected string
	}{
		{name: "1", s: "the sky is blue", expected: "blue is sky the"},
		{name: "2", s: "  hello world  ", expected: "world hello"},
		{name: "3", s: "a good   example", expected: "example good a"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseWords(tt.s)
			want := tt.expected

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}
