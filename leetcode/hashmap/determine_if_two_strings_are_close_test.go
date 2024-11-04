package hashmap

import "testing"

func TestCloseStrings(t *testing.T) {
	testCases := []struct {
		name     string
		word1    string
		word2    string
		expected bool
	}{
		{name: "1", word1: "abc", word2: "bca", expected: true},
		{name: "2", word1: "a", word2: "aa", expected: false},
		{name: "3", word1: "cabbba", word2: "abbccc", expected: true},
		{name: "4", word1: "abbbzcf", word2: "babzzcz", expected: false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := CloseStrings(tt.word1, tt.word2)
			want := tt.expected

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
