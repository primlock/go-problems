package leetcode

import "testing"

func TestValidPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected bool
	}{
		{name: "1", s: "A man, a plan, a canal: Panama", expected: true},
		{name: "2", s: "race a car", expected: false},
		{name: "3", s: " ", expected: true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidPalindrome(tt.s)
			want := tt.expected

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
