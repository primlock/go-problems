package array

import "testing"

func TestLongestConsecutiveSequence(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{name: "1", nums: []int{100, 4, 200, 1, 3, 2}, expected: 4},
		{name: "2", nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, expected: 9},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestConsecutiveSequence(tt.nums)
			want := tt.expected

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}
