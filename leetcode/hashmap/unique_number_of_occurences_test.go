package hashmap

import "testing"

func TestUniqueOccurrences(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expected bool
	}{
		{name: "1", arr: []int{1, 2, 2, 1, 1, 3}, expected: true},
		{name: "2", arr: []int{1, 2}, expected: false},
		{name: "3", arr: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, expected: true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := UniqueOccurrences(tt.arr)
			want := tt.expected

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
