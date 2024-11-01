package playground

import "testing"

func TestPrefixSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		i        int
		j        int
		expected int
	}{
		{name: "1", nums: []int{1, 2, 3, 4, 5}, i: 1, j: 3, expected: 9},
		{name: "2", nums: []int{2, 4, 5, 7, 9}, i: 2, j: 4, expected: 21},
		{name: "3", nums: []int{44, 36, 12, 72, 99, 3}, i: 0, j: 5, expected: 266},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := PrefixSum(tt.nums, tt.i, tt.j)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
