package twopointer

import "testing"

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		val      int
		expected int
	}{
		{name: "1", nums: []int{3, 2, 2, 3}, val: 3, expected: 2},
		{name: "2", nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, expected: 5},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveElement(tt.nums, tt.val)
			want := tt.expected

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}
