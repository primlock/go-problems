package linkedlist

import "testing"

func TestPairSum(t *testing.T) {
	testCases := []struct {
		name     string
		head     *ListNode
		expected int
	}{
		{
			name:     "1",
			head:     &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}},
			expected: 6,
		},
		{
			name:     "2",
			head:     &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}},
			expected: 7,
		},
		{
			name:     "3",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 100000, Next: nil}},
			expected: 100001,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := PairSum(tt.head)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
