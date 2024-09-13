package twopointer

import "testing"

func TestIsSubsequence(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{name: "1", s: "abc", t: "ahbgdc", expected: true},
		{name: "2", s: "axc", t: "ahbgdc", expected: false},
		{name: "3", s: "", t: "looool", expected: true},
		{name: "4", s: "a", t: "", expected: false},
		{name: "5", s: "", t: "", expected: true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSubsequence(tt.s, tt.t)
			want := tt.expected

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
