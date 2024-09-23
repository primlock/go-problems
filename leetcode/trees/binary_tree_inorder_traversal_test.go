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
		{
			name: "1",
			root: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil},
					Right: nil},
			},
			expected: []int{1, 3, 2},
		},
		{name: "2", root: nil, expected: []int{}},
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
