package slidingwindow

import "testing"

func TestLongestSubarray(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{name: "1", nums: []int{1, 1, 0, 1}, expected: 3},
		{name: "2", nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1}, expected: 5},
		{name: "3", nums: []int{1, 1, 1}, expected: 2},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestSubarray(tt.nums)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
