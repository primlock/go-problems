package basic

import "testing"

func TestDiagonalDifference(t *testing.T) {
	testCases := []struct {
		name     string
		arr      [][]int32
		expected int32
	}{
		{name: "1", arr: [][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}, expected: 2},
		{name: "2", arr: [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}, expected: 15},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := DiagonalDifference(tt.arr)
			want := tt.expected

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}
