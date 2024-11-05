package stack

import "testing"

func TestRemoveStars(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected string
	}{
		{name: "1", s: "leet**cod*e", expected: "lecoe"},
		{name: "2", s: "erase*****", expected: ""},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveStars(tt.s)
			want := tt.expected

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}
