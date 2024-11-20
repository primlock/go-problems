package trees

import "testing"

func TestLeafSimilar(t *testing.T) {
	testCases := []struct {
		name     string
		root1    *TreeNode
		root2    *TreeNode
		expected bool
	}{
		{
			name:     "1",
			root1:    ArrayToBinaryTree([]interface{}{3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4}),
			root2:    ArrayToBinaryTree([]interface{}{3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8}),
			expected: true,
		},
		{
			name:     "2",
			root1:    ArrayToBinaryTree([]interface{}{1, 2, 3}),
			root2:    ArrayToBinaryTree([]interface{}{1, 3, 2}),
			expected: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := LeafSimilar(tt.root1, tt.root2)
			want := tt.expected

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
