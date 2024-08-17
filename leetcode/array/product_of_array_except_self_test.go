package array

import (
	"reflect"
	"testing"
)

func TestProductOfArrayExceptSelf(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{name: "1", nums: []int{1, 2, 3, 4}, expected: []int{24, 12, 8, 6}},
		{name: "2", nums: []int{-1, 1, 0, -3, 3}, expected: []int{0, 0, 9, 0, 0}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := ProductOfArrayExceptSelf(tt.nums)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
