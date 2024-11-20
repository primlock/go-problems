package trees

import "testing"

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "1",
			root:     ArrayToBinaryTree([]interface{}{3, 9, 20, nil, nil, 15, 7}),
			expected: 3,
		},
		{
			name:     "2",
			root:     ArrayToBinaryTree([]interface{}{1, nil, 2}),
			expected: 2,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxDepth(tt.root)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
