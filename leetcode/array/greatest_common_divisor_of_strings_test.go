package array

import "testing"

func TestGcdOfStrings(t *testing.T) {
	testCases := []struct {
		name     string
		str1     string
		str2     string
		expected string
	}{
		{name: "1", str1: "ABCABC", str2: "ABC", expected: "ABC"},
		{name: "2", str1: "ABABAB", str2: "ABAB", expected: "AB"},
		{name: "3", str1: "LEET", str2: "CODE", expected: ""},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := GcdOfStrings(tt.str1, tt.str2)
			want := tt.expected

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}
