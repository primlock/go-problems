/*
https://leetcode.com/problems/linked-list-cycle/

There is a cycle in a linked list if there is some node in the list that can be
reached again by continuously following the next pointer.

To solve this problem we are going to create a race between two node pointers.
We create a loop that continues if the fast pointer and fast.next pointers are
not nil. fast advances two pointers at a time (fast.next.next) while slow only
advances one pointer at a time (slow.next). If fast and slow are ever the same
pointer, we have found a cycle.

If there is a cycle there will be no point when fast or fast.next is set to nil,
which means that fast will eventually lap slow.

The edge we need to look out for is the case where head is nil, in which we just
return false.
*/

package linkedlist

func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}
