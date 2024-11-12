package trees

func GoodNodes(root *TreeNode) int {
	// Good = In the path root to node, there are no nodes with a value greater than x
	var dfs func(root *TreeNode, max int) int
	dfs = func(root *TreeNode, max int) int {
		if root == nil {
			return 0
		}

		good := 0

		// Check if this node is good.
		if root.Val >= max {
			good = 1
		}

		// Check for a new max.
		if root.Val > max {
			max = root.Val
		}

		// Recursivly look for good in left and right subtrees.
		return good + dfs(root.Left, max) + dfs(root.Right, max)
	}

	return dfs(root, root.Val)
}
