// https://leetcode.com/problems/count-complete-tree-nodes

package trees

func CountNodes(root *TreeNode) int {
	// Perform in-order traversal to reach all the nodes.
	count := 0

	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		// Visit the left subtree
		inOrder(root.Left)

		count++

		// Visit the right subtree
		inOrder(root.Right)
	}

	inOrder(root)

	return count
}
