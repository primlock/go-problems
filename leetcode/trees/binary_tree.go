package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ArrayToBST(nodes []interface{}) *TreeNode {
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
