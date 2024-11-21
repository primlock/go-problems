package array

import "testing"

func TestThirdMax(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{name: "1", nums: []int{3, 2, 1}, expected: 1},
		{name: "2", nums: []int{2, 1}, expected: 2},
		{name: "3", nums: []int{2, 2, 3, 1}, expected: 1},
		{name: "4", nums: []int{3, 2, 1}, expected: 1},
		{name: "5", nums: []int{5, 2, 2}, expected: 5},
		{name: "6", nums: []int{1, 2, 2, 5, 3, 5}, expected: 2},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := ThirdMax(tt.nums)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
