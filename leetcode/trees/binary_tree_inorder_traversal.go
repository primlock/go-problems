// https://leetcode.com/problems/binary-tree-inorder-traversal/

package trees

func InorderTraversal(root *TreeNode) []int {
	values := make([]int, 0)

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		values = append(values, node.Val)
		dfs(node.Right)
	}

	dfs(root)

	return values
}
