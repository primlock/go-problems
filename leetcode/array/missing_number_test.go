package array

import "testing"

func TestMissingNumber(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{name: "1", nums: []int{3, 0, 1}, expected: 2},
		{name: "2", nums: []int{0, 1}, expected: 2},
		{name: "3", nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, expected: 8},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := MissingNumber(tt.nums)
			want := tt.expected

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
