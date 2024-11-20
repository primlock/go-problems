package trees

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		key      int
		expected *TreeNode
	}{
		{
			name:     "1",
			root:     ArrayToBinaryTree([]interface{}{5, 3, 6, 2, 4, nil, 7}),
			key:      3,
			expected: ArrayToBinaryTree([]interface{}{5, 4, 6, 2, nil, nil, 7})},
		{
			name:     "2",
			root:     ArrayToBinaryTree([]interface{}{5, 3, 6, 2, 4, nil, 7}),
			key:      0,
			expected: ArrayToBinaryTree([]interface{}{5, 3, 6, 2, 4, nil, 7}),
		},
		{
			name:     "3",
			root:     nil,
			key:      0,
			expected: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := DeleteNode(tt.root, tt.key)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
