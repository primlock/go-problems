/*
https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/

Give leader an n head start over follower and then advance both pointers together
until leader is nil. This will place follower at the node before the nth node
making it easy to reassign the pointer to the node that follows the nth node.

We use a dummy node to start our pointers before the head incase it's the head that needs
to be removed.
*/

package linkedlist

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	// Advance the leader to be n nodes ahead
	leader := dummy
	for i := 0; i <= n; i++ {
		leader = leader.Next
	}

	// Advance both together to maintain the gap
	follower := dummy
	for leader != nil {
		leader = leader.Next
		follower = follower.Next
	}

	// Remove the nth node from the end
	follower.Next = follower.Next.Next

	return dummy.Next
}
