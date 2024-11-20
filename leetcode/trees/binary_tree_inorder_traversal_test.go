package trees

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{name: "1", root: ArrayToBinaryTree([]interface{}{1, nil, 2, 3}), expected: []int{1, 3, 2}},
		{name: "2", root: ArrayToBinaryTree([]interface{}{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9}), expected: []int{4, 2, 6, 5, 7, 1, 3, 9, 8}},
		{name: "3", root: nil, expected: []int{}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := InorderTraversal(tt.root)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
