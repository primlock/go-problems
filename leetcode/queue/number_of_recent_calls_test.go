package queue

import "testing"

func TestRecentCounter(t *testing.T) {
	testCases := []struct {
		name     string
		t        int
		expected int
	}{
		{name: "1", t: 1, expected: 1},
		{name: "2", t: 100, expected: 2},
		{name: "3", t: 3001, expected: 3},
		{name: "4", t: 3002, expected: 3},
	}

	r := Constructor()

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := r.Ping(tt.t)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
