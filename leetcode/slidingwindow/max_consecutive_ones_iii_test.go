package slidingwindow

import "testing"

func TestLongestOnes(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{name: "1", nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, k: 2, expected: 6},
		{name: "2", nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, k: 3, expected: 10},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestOnes(tt.nums, tt.k)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
