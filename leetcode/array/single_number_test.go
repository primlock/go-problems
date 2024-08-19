package array

import "testing"

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{name: "1", nums: []int{2, 2, 1}, expected: 1},
		{name: "2", nums: []int{4, 1, 2, 1, 2}, expected: 4},
		{name: "3", nums: []int{1}, expected: 1},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := SingleNumber(tt.nums)
			want := tt.expected

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}
