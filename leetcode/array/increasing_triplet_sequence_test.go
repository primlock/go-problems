package array

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{name: "1", nums: []int{1, 2, 3, 4, 5}, expected: true},
		{name: "2", nums: []int{5, 4, 3, 2, 1}, expected: false},
		{name: "3", nums: []int{2, 1, 5, 0, 4, 6}, expected: true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := IncreasingTriplet(tt.nums)
			want := tt.expected

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
