package twopointer

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{name: "1", nums: []int{1, 1, 2}, expected: 2},
		{name: "2", nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, expected: 5},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveDuplicates(tt.nums)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
