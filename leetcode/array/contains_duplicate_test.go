package array

import "testing"

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{name: "1", nums: []int{1, 2, 3, 1}, expected: true},
		{name: "2", nums: []int{1, 2, 3, 4}, expected: false},
		{name: "3", nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, expected: true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsDuplicate(tt.nums)
			want := tt.expected

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
