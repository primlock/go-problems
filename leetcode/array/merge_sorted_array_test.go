package array

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		name     string
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{name: "1", nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, expected: []int{1, 2, 2, 3, 5, 6}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.nums1, tt.m, tt.nums2, tt.n)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})

	}
}
