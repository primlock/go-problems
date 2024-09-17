package linkedlist

func RemoveElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	current := head
	var prev *ListNode

	for current != nil {
		if current.Val == val {
			if current == head {
				head = current.Next
			} else {
				prev.Next = current.Next
			}
			current = current.Next
		} else {
			prev = current
			current = current.Next
		}
	}

	return head
}
