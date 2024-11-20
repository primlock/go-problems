package trees

import "testing"

func TestGoodNodes(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "1",
			root:     ArrayToBinaryTree([]interface{}{3, 1, 4, 3, nil, 1, 5}),
			expected: 4,
		},
		{
			name:     "2",
			root:     ArrayToBinaryTree([]interface{}{3, 3, nil, 4, 2}),
			expected: 3,
		},
		{
			name:     "3",
			root:     ArrayToBinaryTree([]interface{}{2, nil, 4, 10, 8, nil, nil, 4}),
			expected: 4,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := GoodNodes(tt.root)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
