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
			name: "1",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:  6,
					Left: nil,
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			key: 3,
			expected: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val:  6,
					Left: nil,
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
		{
			name:     "2",
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
