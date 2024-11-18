// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree

package array

import (
	"github.com/reyncode/go-problems/leetcode/trees"
)

type TreeNode = trees.TreeNode

func insert(array []int, start, end int) *TreeNode {
	// Guard from overlapping boundaries
	if start > end {
		return nil
	}

	// Find the middle element and turn it into a node
	mid := (start + end) / 2
	root := &TreeNode{Val: array[mid]}

	// Build the left subtree with the left half of the array
	root.Left = insert(array, start, mid-1)

	// Build the right subtree with the right half of the array
	root.Right = insert(array, mid+1, end)

	return root
}

func SortedArrayToBST(nums []int) *TreeNode {
	// We need to build our BST by dividing the input array in half where the left side
	// is in the left subtree and the right side is in the right subtree.

	// Use binary search to determine which element should be inserted next.
	return insert(nums, 0, len(nums)-1)
}
