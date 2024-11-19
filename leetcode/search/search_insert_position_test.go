package search

import "testing"

func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{name: "1", nums: []int{1, 3, 5, 6}, target: 5, expected: 2},
		{name: "2", nums: []int{1, 3, 5, 6}, target: 2, expected: 1},
		{name: "3", nums: []int{1, 3, 5, 6}, target: 7, expected: 4},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchInsert(tt.nums, tt.target)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
