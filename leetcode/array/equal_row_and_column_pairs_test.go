package array

import "testing"

func TestEqualPairs(t *testing.T) {
	testCases := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{name: "1", grid: [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}, expected: 1},
		{name: "2", grid: [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}, expected: 3},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := EqualPairs(tt.grid)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
