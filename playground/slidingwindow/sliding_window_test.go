package playground

import "testing"

func TestMaxSumInSubset(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{name: "1", nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, k: 3, expected: 27},
		{name: "2", nums: []int{1, 12, -5, -6, 50, 3}, k: 4, expected: 51},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxSumInSubset(tt.nums, tt.k)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
