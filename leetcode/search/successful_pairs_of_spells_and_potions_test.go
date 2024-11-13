package search

import (
	"reflect"
	"testing"
)

func TestSuccessfulPairs(t *testing.T) {
	testCases := []struct {
		name     string
		spells   []int
		potions  []int
		success  int64
		expected []int
	}{
		{name: "1", spells: []int{5, 1, 3}, potions: []int{1, 2, 3, 4, 5}, success: 7, expected: []int{4, 0, 3}},
		{name: "2", spells: []int{3, 1, 2}, potions: []int{8, 5, 8}, success: 16, expected: []int{2, 0, 2}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := SuccessfulPairs(tt.spells, tt.potions, tt.success)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
