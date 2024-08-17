package array

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{name: "1", nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
		{name: "2", nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
		{name: "3", nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
