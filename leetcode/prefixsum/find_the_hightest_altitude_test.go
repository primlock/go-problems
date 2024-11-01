package prefixsum

import "testing"

func TestLargestAltitude(t *testing.T) {
	testCases := []struct {
		name     string
		gain     []int
		expected int
	}{
		{name: "1", gain: []int{-5, 1, 5, 0, -7}, expected: 1},
		{name: "2", gain: []int{-4, -3, -2, -1, 4, 3, 2}, expected: 0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := LargestAltitude(tt.gain)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
