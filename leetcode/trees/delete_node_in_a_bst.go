// https://leetcode.com/problems/delete-node-in-a-bst/

package trees

func successor(root *TreeNode) int {
	// Traverse right and then left until you get to the end
	root = root.Right

	for root.Left != nil {
		root = root.Left
	}

	return root.Val
}

func predecessor(root *TreeNode) int {
	root = root.Left

	for root.Right != nil {
		root = root.Right
	}

	return root.Val
}

func DeleteNode(root *TreeNode, key int) *TreeNode {
	// Make sure the key is in the tree
	if root == nil {
		return nil
	}

	// Find the node's location
	if key == root.Val {
		// Check if this is a leaf
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Right != nil {
			// find the successor
			root.Val = successor(root)
			root.Right = DeleteNode(root.Right, root.Val)
		} else {
			// find the predecessor
			root.Val = predecessor(root)
			root.Left = DeleteNode(root.Left, root.Val)
		}
	} else if key < root.Val {
		root.Left = DeleteNode(root.Left, key)
	} else {
		root.Right = DeleteNode(root.Right, key)
	}

	return root
}
