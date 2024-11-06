package stack

import "testing"

func TestDecodeString(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected string
	}{
		{name: "1", s: "3[a]2[bc]", expected: "aaabcbc"},
		{name: "2", s: "3[a2[c]]", expected: "accaccacc"},
		{name: "3", s: "2[abc]3[cd]ef", expected: "abcabccdcdcdef"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := DecodeString(tt.s)
			want := tt.expected

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}
