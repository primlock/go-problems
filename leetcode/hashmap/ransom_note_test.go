package hashmap

import "testing"

func TestCanConstruct(t *testing.T) {
	testCases := []struct {
		name       string
		ransomNote string
		magazine   string
		expected   bool
	}{
		{name: "1", ransomNote: "a", magazine: "b", expected: false},
		{name: "2", ransomNote: "aa", magazine: "ab", expected: false},
		{name: "3", ransomNote: "aa", magazine: "aab", expected: true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := canConstruct(tt.ransomNote, tt.magazine)

			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
