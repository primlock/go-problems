package basic

import "testing"

func TestPlusMinus(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int32
	}{
		{name: "1", arr: []int32{-4, 3, -9, 0, 4, 1}},
		{name: "2", arr: []int32{1, 2, 3, -1, -2, -3, 0, 0}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			PlusMinus(tt.arr)
		})
	}
}
