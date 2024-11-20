package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// This function builds a binary tree from an interface.
// The programmer can pass nil as an element to represent the end
// of a subtree.
// The output of this function is NOT garunteed to be balanced or
// follow the rules of a binary search tree.
func ArrayToBinaryTree(nodes []interface{}) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}

	// Create the root node from the first element
	root := &TreeNode{Val: nodes[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	// Process the queue and insert nodes
	for i < len(nodes) && len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(nodes) && nodes[i] != nil {
			current.Left = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, current.Left)
		}
		i++

		// Right child
		if i < len(nodes) && nodes[i] != nil {
			current.Right = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, current.Right)
		}
		i++
	}

	return root
}
