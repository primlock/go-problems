// https://leetcode.com/problems/reverse-linked-list

package linkedlist

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next

		// Point the current node backwards
		current.Next = prev

		// Advance
		prev = current
		current = next
	}

	return prev
}
