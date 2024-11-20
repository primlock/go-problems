package trees

import (
	"reflect"
	"testing"
)

func TestSearchBST(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		val      int
		expected *TreeNode
	}{
		{
			name:     "1",
			root:     ArrayToBinaryTree([]interface{}{4, 2, 7, 1, 3}),
			val:      2,
			expected: ArrayToBinaryTree([]interface{}{2, 1, 3}),
		},
		{
			name:     "2",
			root:     ArrayToBinaryTree([]interface{}{4, 2, 7, 1, 3}),
			val:      5,
			expected: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchBST(tt.root, tt.val)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
