package array

import "testing"

func TestMakeFancyString(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected string
	}{
		{name: "1", s: "leeetcode", expected: "leetcode"},
		{name: "2", s: "aaabaaaa", expected: "aabaa"},
		{name: "3", s: "aab", expected: "aab"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := MakeFancyString(tt.s)
			want := tt.expected

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}
