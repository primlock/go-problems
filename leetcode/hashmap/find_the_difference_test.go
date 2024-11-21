package hashmap

import "testing"

func TestFindTheDifference(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		t        string
		expected byte
	}{
		{name: "1", s: "abcd", t: "abcde", expected: 'e'},
		{name: "2", s: "", t: "y", expected: 'y'},
		{name: "3", s: "a", t: "aa", expected: 'a'},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := FindTheDifference(tt.s, tt.t)
			want := tt.expected

			if got != want {
				t.Errorf("got %c, want %c", got, want)
			}
		})
	}
}
