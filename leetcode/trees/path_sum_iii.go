package trees

func PathSum(root *TreeNode, targetSum int) int {
	// Base case for a path starting from root
	prefixSumCount := map[int]int{0: 1}
	return dfs(root, 0, targetSum, prefixSumCount)
}

func dfs(node *TreeNode, currentSum, targetSum int, prefixSumCount map[int]int) int {
	if node == nil {
		return 0
	}

	currentSum += node.Val

	// Check if there's a valid path to target
	paths := prefixSumCount[currentSum-targetSum]

	// Update the prefix sum map with the current sum
	prefixSumCount[currentSum]++

	// Recurse on left and right subtrees
	paths += dfs(node.Left, currentSum, targetSum, prefixSumCount)
	paths += dfs(node.Right, currentSum, targetSum, prefixSumCount)

	// Backtrack: remove the current sum from the map to not interfere with other paths
	prefixSumCount[currentSum]--

	return paths
}
