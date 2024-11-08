// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list

package linkedlist

func PairSum(head *ListNode) int {
	// Map the nodes while obtaining n for a quick lookup of node Val. (k: i, v: Val)
	nodeMap := map[int]int{}
	n := 0
	max := 0

	// Find n and fill the map
	current := head
	for current != nil {
		nodeMap[n] = current.Val
		n++
		current = current.Next
	}

	// Lookup the values for the twin and compare to max.
	current = head
	i := 0
	for current != nil {
		// Find this nodes twin and what the sum is.
		twinIndex := n - 1 - i
		twinSum := current.Val + nodeMap[twinIndex]

		if twinSum > max {
			max = twinSum
		}

		i++
		current = current.Next
	}

	return max
}
