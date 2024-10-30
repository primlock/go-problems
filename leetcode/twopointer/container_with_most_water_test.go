package twopointer

import "testing"

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		name     string
		height   []int
		expected int
	}{
		{name: "1", height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, expected: 49},
		{name: "2", height: []int{1, 1}, expected: 1},
		{name: "3", height: []int{8, 7, 2, 1}, expected: 7},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxArea(tt.height)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
