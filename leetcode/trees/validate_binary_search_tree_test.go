package trees

import "testing"

func TestIsValidBST(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name:     "1",
			root:     ArrayToBinaryTree([]interface{}{2, 1, 3}),
			expected: true,
		},
		{
			name:     "2",
			root:     ArrayToBinaryTree([]interface{}{5, 1, 4, nil, nil, 3, 6}),
			expected: false,
		},
		{
			name:     "3",
			root:     ArrayToBinaryTree([]interface{}{5, 4, 6, nil, nil, 3, 7}),
			expected: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidBST(tt.root)
			want := tt.expected

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
