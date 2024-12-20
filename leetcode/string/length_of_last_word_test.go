package lstring

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{name: "1", s: "Hello World", expected: 5},
		{name: "2", s: "   fly me   to   the moon  ", expected: 4},
		{name: "3", s: "luffy is still joyboy", expected: 6},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLastWord(tt.s)

			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
