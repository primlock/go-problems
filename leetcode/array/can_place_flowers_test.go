package array

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	testCases := []struct {
		name      string
		flowerbed []int
		n         int
		expected  bool
	}{
		{name: "1", flowerbed: []int{0, 0, 1, 0, 0}, n: 1, expected: true},
		{name: "2", flowerbed: []int{1, 0, 0, 0, 1}, n: 1, expected: true},
		{name: "3", flowerbed: []int{1, 0, 0, 0, 1}, n: 2, expected: false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := CanPlaceFlowers(tt.flowerbed, tt.n)
			want := tt.expected

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
