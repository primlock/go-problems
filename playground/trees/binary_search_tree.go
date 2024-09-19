/*
Binary Search Tree Order:

All nodes to the left of the current node are less than equal to the current
node while all nodes to the right of the current node are greater than the
current node.

	2
   / \
  1   3

*/

package playground

import "fmt"

// Insert - Decide if the values needs to be inserted to the left or the right of the root.
func (n *TreeNode) Insert(data int) (node *TreeNode) {
	if n.Search(data) {
		return nil
	}

	if n == nil {
		return &TreeNode{data: data, left: nil, right: nil}
	} else if n.data > data { // left
		// Determine if we need to keep going deeper
		if n.left == nil {
			n.left = &TreeNode{data: data, left: nil, right: nil}
		} else {
			// Need to go further to find a vacancy
			n.left.Insert(data)
		}
	} else { // right
		if n.right == nil {
			n.right = &TreeNode{data: data, left: nil, right: nil}
		} else {
			n.right.Insert(data)
		}
	}

	return n
}

// Search - Looking through the tree to find a target value
func (n *TreeNode) Search(target int) bool {
	if n == nil {
		return false
	} else {
		if n.data == target {
			return true
		} else if target < n.data { // left
			return n.left.Search(target)
		} else { // right
			return n.right.Search(target)
		}
	}
}

/*
Traversal - Visiting all the nodes of the tree

There are 4 main algorithms to travel through the tree:
- In-Order: Visits left subtree, then the node, then the right subtree.
- Pre-Order: Visits the node first, then left subtree, then right subtree.
- Post-Order: Visits left subtree, then right subtree, then the node.
- Level-Order: Visits all nodes at the present depth before moving on to nodes
			   at the next depth level. It is typically implemented using a queue.

Depth First Search (DFS) algorithms (in, pre, post) explore down a branch as far as possible before
backtracking.

Breadth First Search (BFS) algorithms (level) explores the node at the current level before
going to the next level.
*/

// In-order (left-subtree, root, right-subtree)
func (n *TreeNode) InOrderTraversal() []int {
	array := make([]int, 0)

	var inOrder func(*TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inOrder(root.left)
		array = append(array, root.data)
		inOrder(root.right)
	}

	inOrder(n)

	return array
}

// Pre-order (root, left-subtree, right-subtree)
func (n *TreeNode) PreOrderTraversal() {
	if n == nil {
		return
	}

	fmt.Println(n.data)
	n.left.PreOrderTraversal()
	n.right.PreOrderTraversal()
}

// Post-order (left-subtree, right-subtree, root)
func (n *TreeNode) PostOrderTraversal() {
	if n == nil {
		return
	}

	n.left.PostOrderTraversal()
	n.right.PostOrderTraversal()
	fmt.Println(n.data)
}

// Level-order (current level, next level, ...)
func (n *TreeNode) LevelOrderTraversal() {
	if n == nil {
		return
	}

	queue := []*TreeNode{n}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node.data)

		if node.left != nil {
			queue = append(queue, node.left)
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}

/*
Deletion - Removing a value from the BST

There are 3 situations we need to address:
- The node we want to remove is a leaf
- The node we want to remove has one child
- The node we want to remove has two children

To keep our BST a BST after removing a node we need to make sure we
assign the proper value in place of the removed node. We have to
accurately select a successor or a predecessor.

We can make this determination by looking to see if a node exists to the
right of the root node. If it does, we need to search for a successor. If
it doesn't, we need to look for a predecessor.
*/

func (n *TreeNode) Remove(data int) *TreeNode {
	if n == nil {
		return nil
	} else if !n.Search(data) {
		return nil
	}

	// Find the node's location
	if data == n.data {
		// Is this a leaf?
		if n.left == nil && n.right == nil {
			n = nil
		} else if n.right != nil {
			// need a successor
			n.data = n.Successor()
			n.right = n.right.Remove(data)
		} else {
			// need a predecessor
			n.data = n.Predecessor()
			n.left = n.left.Remove(data)
		}
	} else if data < n.data {
		n.left = n.left.Remove(data)
	} else {
		n.right = n.right.Remove(data)
	}

	// Return the new root
	return n
}

func (n *TreeNode) Successor() int {
	// Go right and then left until you hit the end
	n = n.right

	for n.left != nil {
		n = n.left
	}

	return n.data
}

func (n *TreeNode) Predecessor() int {
	// Go left and then right until you hit the end
	n = n.left

	for n.right != nil {
		n = n.right
	}

	return n.data
}
