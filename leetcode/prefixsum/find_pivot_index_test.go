package prefixsum

import "testing"

func TestPivotIndex(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{name: "1", nums: []int{1, 7, 3, 6, 5, 6}, expected: 3},
		{name: "2", nums: []int{1, 2, 3}, expected: -1},
		{name: "3", nums: []int{2, 1, -1}, expected: 0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := PivotIndex(tt.nums)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
