package stack

import (
	"reflect"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	testCases := []struct {
		name      string
		asteroids []int
		expected  []int
	}{
		{name: "1", asteroids: []int{5, 10, -5}, expected: []int{5, 10}},
		{name: "2", asteroids: []int{8, -8}, expected: []int{}},
		{name: "3", asteroids: []int{10, 2, -5}, expected: []int{10}},
		{name: "4", asteroids: []int{-2, -1, 1, 2}, expected: []int{-2, -1, 1, 2}},
		{name: "5", asteroids: []int{-2, -2, 1, -2}, expected: []int{-2, -2, -2}},
		{name: "6", asteroids: []int{-2, -2, -2, -2}, expected: []int{-2, -2, -2, -2}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := AsteroidCollision(tt.asteroids)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
