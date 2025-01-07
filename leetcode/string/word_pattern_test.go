package lstring

import "testing"

func TestWordPattern(t *testing.T) {
	testCases := []struct {
		name     string
		pattern  string
		s        string
		expected bool
	}{
		{name: "1", pattern: "abba", s: "dog cat cat dog", expected: true},
		{name: "2", pattern: "abba", s: "dog cat cat fish", expected: false},
		{name: "3", pattern: "aaaa", s: "dog cat cat dog", expected: false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := WordPattern(tt.pattern, tt.s)
			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
