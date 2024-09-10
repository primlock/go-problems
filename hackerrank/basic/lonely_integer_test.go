package basic

import "testing"

func TestLonelyInteger(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int32
		expected int32
	}{
		{name: "1", arr: []int32{1, 2, 3, 4, 3, 2, 1}, expected: 4},
		{name: "2", arr: []int32{1}, expected: 1},
		{name: "3", arr: []int32{1, 1, 2}, expected: 2},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := LonelyInteger(tt.arr)
			want := tt.expected

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}
