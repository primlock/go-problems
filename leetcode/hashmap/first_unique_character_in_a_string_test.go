package hashmap

import "testing"

func TestFirstUniqueChar(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{name: "1", s: "leetcode", expected: 0},
		{name: "2", s: "loveleetcode", expected: 2},
		{name: "3", s: "aabb", expected: -1},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := firstUniqueChar(tt.s)

			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
