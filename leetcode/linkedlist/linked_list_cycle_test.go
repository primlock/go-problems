package linkedlist

import "testing"

func TestHasCycle(t *testing.T) {
	testCases := []struct {
		name     string
		pos      int
		head     *ListNode
		expected bool
	}{
		{name: "1", pos: 1, head: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4, Next: nil}}}}, expected: true},
		{name: "2", pos: 0, head: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: nil}}, expected: true},
		{name: "3", pos: -1, head: &ListNode{Val: 3, Next: nil}, expected: false},
	}

	for _, tt := range testCases {
		// Setup the cycle node
		if tt.pos != -1 {
			var counter int
			var node *ListNode

			current := tt.head
			for current.Next != nil {
				if counter == tt.pos {
					node = current
				}
				counter++
				current = current.Next
			}

			current.Next = node
		}

		t.Run(tt.name, func(t *testing.T) {
			got := HasCycle(tt.head)
			want := tt.expected

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
