package array

import "testing"

func TestMergeAlternately(t *testing.T) {
	testCases := []struct {
		name     string
		word1    string
		word2    string
		expected string
	}{
		{name: "1", word1: "abc", word2: "pqr", expected: "apbqcr"},
		{name: "2", word1: "ab", word2: "pqrs", expected: "apbqrs"},
		{name: "3", word1: "abcd", word2: "pq", expected: "apbqcd"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeAlternately(tt.word1, tt.word2)
			want := tt.expected

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}
