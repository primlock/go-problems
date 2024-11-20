package array

import (
	"reflect"
	"testing"

	"github.com/reyncode/go-problems/leetcode/trees"
)

func TestSortedArrayToBST(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected *TreeNode
	}{
		{name: "1", nums: []int{-10, -3, 0, 5, 9}, expected: trees.ArrayToBinaryTree([]interface{}{0, -10, 5, nil, -3, nil, 9})},
		{name: "2", nums: []int{1, 3}, expected: trees.ArrayToBinaryTree([]interface{}{1, nil, 3})},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := SortedArrayToBST(tt.nums)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
