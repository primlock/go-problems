package array

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	testCases := []struct {
		name         string
		candies      []int
		extraCandies int
		expected     []bool
	}{
		{name: "1", candies: []int{2, 3, 5, 1, 3}, extraCandies: 3, expected: []bool{true, true, true, false, true}},
		{name: "2", candies: []int{4, 2, 1, 1, 2}, extraCandies: 1, expected: []bool{true, false, false, false, false}},
		{name: "3", candies: []int{12, 1, 12}, extraCandies: 10, expected: []bool{true, false, true}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := KidsWithCandies(tt.candies, tt.extraCandies)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
