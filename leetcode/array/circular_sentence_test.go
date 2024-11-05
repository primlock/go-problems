package array

import "testing"

func TestIsCircularSentence(t *testing.T) {
	testCases := []struct {
		name     string
		sentence string
		expected bool
	}{
		{name: "1", sentence: "leetcode exercises sound delightful", expected: true},
		{name: "2", sentence: "eetcode", expected: true},
		{name: "3", sentence: "leetcode", expected: false},
		{name: "4", sentence: "Leetcode is cool", expected: false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := IsCircularSentence(tt.sentence)
			want := tt.expected

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
