// https://leetcode.com/problems/maximum-depth-of-binary-tree

package trees

import "math"

func MaxDepth(root *TreeNode) int {
	// Empty Tree.
	if root == nil {
		return 0
	}

	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)

	return int(math.Max(float64(left), float64(right))) + 1
}
