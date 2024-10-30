package twopointer

import "testing"

func TestMaxOperations(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{name: "1", nums: []int{1, 2, 3, 4}, k: 5, expected: 2},
		{name: "2", nums: []int{3, 1, 3, 4, 3}, k: 6, expected: 1},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxOperations(tt.nums, tt.k)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
