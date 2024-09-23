// https://leetcode.com/problems/search-in-a-binary-search-tree/description/

package trees

func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val == root.Val {
		return root
	} else if val < root.Val {
		return SearchBST(root.Left, val)
	} else {
		return SearchBST(root.Right, val)
	}
}
