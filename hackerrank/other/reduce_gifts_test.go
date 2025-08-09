package other

import "testing"

func TestReduceGifts(t *testing.T) {
	testCases := []struct {
		name      string
		prices    []int
		k         int
		threshold int
		expected  int
	}{
		{name: "1", prices: []int{3, 2, 1, 4, 6, 5}, k: 3, threshold: 14, expected: 1},
		{name: "2", prices: []int{1, 2, 3}, k: 4, threshold: 5, expected: 0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := ReduceGifts(tt.prices, tt.k, tt.threshold)
			want := tt.expected

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
