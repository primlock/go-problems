package linkedlist

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	testCases := []struct {
		name     string
		head     *ListNode
		n        int
		expected *ListNode
	}{
		{
			name:     "1",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
			n:        2,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}},
		},
		{
			name:     "2",
			head:     &ListNode{Val: 1, Next: nil},
			n:        1,
			expected: nil,
		},
		{
			name:     "3",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
			n:        1,
			expected: &ListNode{Val: 1, Next: nil},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveNthFromEnd(tt.head, tt.n)
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
