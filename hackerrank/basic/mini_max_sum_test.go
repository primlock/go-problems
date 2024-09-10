package basic

import "testing"

func TestMiniMaxSum(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int32
	}{
		{name: "1", arr: []int32{1, 2, 3, 4, 5}},       // 10 14
		{name: "2", arr: []int32{7, 69, 2, 221, 8974}}, // 299 9271
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			MiniMaxSum(tt.arr)
		})
	}
}
