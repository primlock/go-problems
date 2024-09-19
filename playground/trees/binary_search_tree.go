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

func (t *BinaryTree) BSTInsert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &TreeNode{data: data, left: nil, right: nil}
	} else {
		t.root.BSTInsert(data)
	}

	return t
}

func (n *TreeNode) BSTInsert(data int) {
	if n == nil {
		return
	} else if n.data >= data { // left
		// Determine if we need to keep going deeper
		if n.left == nil {
			n.left = &TreeNode{data: data, left: nil, right: nil}
		} else {
			// Need to go further to find a vacancy
			n.left.BSTInsert(data)
		}
	} else { // right
		if n.right == nil {
			n.right = &TreeNode{data: data, left: nil, right: nil}
		} else {
			n.right.BSTInsert(data)
		}
	}
}

// Search - Looking through the tree to find a target value
func (t *BinaryTree) BSTSearch(target int) bool {
	if t.root == nil {
		return false
	} else {
		return t.root.BSTSearch(target)
	}
}

func (n *TreeNode) BSTSearch(target int) bool {
	if n == nil {
		return false
	} else {
		if n.data == target {
			return true
		} else if target < n.data { // left
			return n.left.BSTSearch(target)
		} else { // right
			return n.right.BSTSearch(target)
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
func (t *BinaryTree) BSTTraversalInOrder() {
	if t.root == nil {
		return
	} else {
		t.root.BSTTraversalInOrder()
	}
}

func (n *TreeNode) BSTTraversalInOrder() {
	if n == nil {
		return
	}

	// Go left as far as possible, then look at node, then go right as far as possible
	n.left.BSTTraversalInOrder()
	fmt.Println(n.data)
	n.right.BSTTraversalInOrder()
}

// Pre-order (root, left-subtree, right-subtree)
func (t *BinaryTree) BSTTraveralPreOrder() {
	if t.root == nil {
		return
	} else {
		t.root.BSTTraversalPreOrder()
	}
}

func (n *TreeNode) BSTTraversalPreOrder() {
	if n == nil {
		return
	} else {
		fmt.Println(n.data)
		n.left.BSTTraversalPreOrder()
		n.right.BSTTraversalPreOrder()
	}
}

// Post-order (left-subtree, right-subtree, root)
func (t *BinaryTree) BSTTraversalPostOrder() {
	if t.root == nil {
		return
	} else {
		t.root.BSTTraversalPostOrder()
	}
}

func (n *TreeNode) BSTTraversalPostOrder() {
	if n == nil {
		return
	} else {
		n.left.BSTTraversalPostOrder()
		n.right.BSTTraversalPostOrder()
		fmt.Println(n.data)
	}
}

// Level-order (current level, next level, ...)
func (t *BinaryTree) BSTTraversalLevelOrder() {
	if t.root == nil {
		return
	} else {
		t.root.BSTTraversalLevelOrder()
	}
}

func (n *TreeNode) BSTTraversalLevelOrder() {
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

func (b *BinaryTree) Remove(data int) {
	// Node must exist in tree to be removed
	if b.BSTSearch(data) {
		Remove(b.root, data)
	}
}

func Remove(root *TreeNode, data int) *TreeNode {
	if root == nil {
		return root
	}

	// Find the node's location
	if data == root.data {
		// Is this a leaf?
		if root.left == nil && root.right == nil {
			root = nil
		} else if root.right != nil {
			// need a successor
			root.data = Successor(root)
			root.right = Remove(root.right, root.data)
		} else {
			// need a predecessor
			root.data = Predecessor(root)
			root.left = Remove(root.left, root.data)
		}
	} else if data < root.data {
		root.left = Remove(root.left, data)
	} else {
		root.right = Remove(root.right, data)
	}

	return root
}

func Successor(root *TreeNode) int {
	// Go right and then left until you hit the end
	root = root.right

	for root.left != nil {
		root = root.left
	}

	return root.data
}

func Predecessor(root *TreeNode) int {
	// Go left and then right until you hit the end
	root = root.left

	for root.right != nil {
		root = root.right
	}

	return root.data
}
