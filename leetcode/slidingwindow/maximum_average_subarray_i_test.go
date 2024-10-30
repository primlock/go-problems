package slidingwindow

import "testing"

func TestFindMaxAverage(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected float64
	}{
		{name: "1", nums: []int{1, 12, -5, -6, 50, 3}, k: 4, expected: 12.75000},
		{name: "2", nums: []int{5}, k: 1, expected: 5.00000},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMaxAverage(tt.nums, tt.k)
			want := tt.expected

			if got != want {
				t.Errorf("got %.5f, want %.5f", got, want)
			}
		})
	}
}
