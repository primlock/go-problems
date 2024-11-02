package hashmap

import (
	"reflect"
	"testing"
)

func TestFindDifference(t *testing.T) {
	testCases := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected [][]int
	}{
		{name: "1", nums1: []int{1, 2, 3}, nums2: []int{2, 4, 6}, expected: [][]int{{1, 3}, {4, 6}}},
		{name: "2", nums1: []int{1, 2, 3, 3}, nums2: []int{1, 1, 2, 2}, expected: [][]int{{3}, nil}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := FindDifference(tt.nums1, tt.nums2)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
