// https://leetcode.com/problems/odd-even-linked-list

package linkedlist

func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}

	// Link the end of odd list to the head of even list.
	odd.Next = evenHead

	return head
}
