package playground

import "testing"

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{name: "1", arr: []int{2, 3, 4, 10, 40}, target: 3, expected: 1},
		{name: "2", arr: []int{6, 7, 23, 56, 77}, target: 19, expected: -1},
		{name: "3", arr: []int{}, target: 7, expected: -1},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.arr, 0, len(tt.arr)-1, tt.target)
			want := tt.expected

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}
