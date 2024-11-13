// https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree

package trees

func LongestZigZag(root *TreeNode) int {
	// zig-zag is going right node then left node over and over until you can't go any further.
	// length is the number of nodes visited - 1.

	max := 0

	var dfs func(node *TreeNode, isLeft bool, length int)
	dfs = func(node *TreeNode, isLeft bool, length int) {
		if node == nil {
			return
		}

		// Check if we have reached a new max length
		if length > max {
			max = length
		}

		// Traverse in a zig-zag pattern
		if isLeft {
			// Continue the path, go right this time (left = false) and increment length
			dfs(node.Left, false, length+1)
			// Start a new zig-zag path, reset length
			dfs(node.Right, true, 1)
		} else {
			// Continue the path, go left this time (left = true) and increment length
			dfs(node.Right, true, length+1)
			// Start a new zig-zag path, reset length
			dfs(node.Left, false, 1)
		}
	}

	// Run dfs for left and right subtrees
	dfs(root, true, 0)
	dfs(root, false, 0)

	return max
}
