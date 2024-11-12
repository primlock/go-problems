package trees

import "testing"

func TestPathSum(t *testing.T) {
	testCases := []struct {
		name      string
		root      *TreeNode
		targetSum int
		expected  int
	}{
		{name: "1", root: ArrayToBST([]interface{}{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1}), targetSum: 8, expected: 3},
		{name: "2", root: ArrayToBST([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1}), targetSum: 22, expected: 3},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := PathSum(tt.root, tt.targetSum)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
