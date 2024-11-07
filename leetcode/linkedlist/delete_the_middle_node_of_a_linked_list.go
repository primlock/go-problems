package linkedlist

func DeleteMiddle(head *ListNode) *ListNode {
	// Find n, the size of the linked list.
	n := 0
	current := head
	for current != nil {
		n++
		current = current.Next
	}

	// Delete the only node in the list if size is 1.
	if n == 1 {
		return nil
	}

	mid := n / 2
	current = head
	var prev *ListNode

	// Traverse to mid and connect previous next to current's next.
	for i := 0; i < n; i++ {
		if i == mid {
			prev.Next = current.Next
			break
		}

		prev = current
		current = current.Next
	}

	return head
}
