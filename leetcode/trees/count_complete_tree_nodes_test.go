package trees

import "testing"

func TestCountNodes(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{name: "1", root: ArrayToBinaryTree([]interface{}{1, 2, 3, 4, 5, 6}), expected: 6},
		{name: "2", root: nil, expected: 0},
		{name: "3", root: ArrayToBinaryTree([]interface{}{1}), expected: 1},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := CountNodes(tt.root)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
