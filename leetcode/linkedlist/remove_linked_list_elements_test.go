package linkedlist

import "testing"

func TestRemoveElements(t *testing.T) {
	testCases := []struct {
		name     string
		head     *ListNode
		val      int
		expected *ListNode
	}{
		{name: "1", head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}}, val: 6, expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}},
		{name: "2", head: nil, val: 6, expected: nil},
		{name: "3", head: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: nil}}}}, val: 7, expected: nil},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveElements(tt.head, tt.val)
			want := tt.expected

			for got != nil {
				if got.Val != want.Val {
					t.Errorf("got %d want %d", got.Val, want.Val)
				}
				got = got.Next
				want = want.Next
			}
		})
	}
}
