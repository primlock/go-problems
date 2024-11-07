package stack

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected bool
	}{
		{name: "1", s: "()", expected: true},
		{name: "2", s: "()[]{}", expected: true},
		{name: "3", s: "(]", expected: false},
		{name: "4", s: "([])", expected: true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValid(tt.s)
			want := tt.expected

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
