// https://leetcode.com/problems/leaf-similar-trees

package trees

import "slices"

func findLeaves(root *TreeNode, sequence *[]int) {
	if root == nil {
		return
	}

	// To find a leaf, node is not nil but left and right are both nil
	if root.Left == nil && root.Right == nil {
		*sequence = append(*sequence, root.Val)
		return
	}

	findLeaves(root.Left, sequence)
	findLeaves(root.Right, sequence)
}

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafValueSequenceOne := []int{}
	leafValueSequenceTwo := []int{}

	findLeaves(root1, &leafValueSequenceOne)
	findLeaves(root2, &leafValueSequenceTwo)

	// Compare both lists for equality
	return slices.Equal(leafValueSequenceOne, leafValueSequenceTwo)
}
