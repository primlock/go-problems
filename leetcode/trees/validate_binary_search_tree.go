// https://leetcode.com/problems/validate-binary-search-tree/description/

package trees

import "math"

func IsValidBST(root *TreeNode) bool {
	// Only one node
	if root.Left == nil && root.Right == nil {
		return true
	}

	// Do BFS - level order traversal
	queue := []struct {
		node *TreeNode
		min  int
		max  int
	}{
		{root, math.MinInt64, math.MaxInt64},
	}

	for len(queue) > 0 {
		// Pop a node from the queue
		current := queue[0]
		queue = queue[1:]

		node, min, max := current.node, current.min, current.max

		// Make sure the current node is within the allowed range
		if node.Val <= min || node.Val >= max {
			return false
		}

		// If there is a left child, its value must be less than the current node's value
		if node.Left != nil {
			queue = append(queue, struct {
				node     *TreeNode
				min, max int
			}{node.Left, min, node.Val})
		}

		// If there is a right child, its value must be greater than the current node's value
		if node.Right != nil {
			queue = append(queue, struct {
				node     *TreeNode
				min, max int
			}{node.Right, node.Val, max})
		}
	}

	return true
}
