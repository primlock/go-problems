package trees

import "testing"

func TestLongestZigZag(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{name: "1", root: ArrayToBST([]interface{}{1, nil, 1, 1, 1, nil, nil, 1, 1, nil, 1, nil, nil, nil, 1}), expected: 3},
		{name: "2", root: ArrayToBST([]interface{}{1, 1, 1, nil, 1, nil, nil, 1, 1, nil, 1}), expected: 4},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestZigZag(tt.root)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
