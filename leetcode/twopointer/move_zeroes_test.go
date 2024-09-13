package twopointer

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{name: "1", nums: []int{0, 1, 0, 3, 12}, expected: []int{1, 3, 12, 0, 0}},
		{name: "2", nums: []int{0}, expected: []int{0}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := MoveZeroes(tt.nums)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
