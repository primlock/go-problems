package search

import "testing"

func TestGuessNumber(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		pick     int
		expected int
	}{
		{name: "1", n: 10, pick: 6, expected: 6},
		{name: "2", n: 1, pick: 1, expected: 1},
		{name: "3", n: 2, pick: 1, expected: 1},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := GuessNumber(tt.n, tt.pick)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
