package array

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	testCases := []struct {
		name     string
		digits   []int
		expected []int
	}{
		{name: "1", digits: []int{1, 2, 3}, expected: []int{1, 2, 4}},
		{name: "2", digits: []int{4, 3, 2, 1}, expected: []int{4, 3, 2, 2}},
		{name: "3", digits: []int{9}, expected: []int{1, 0}},
		{name: "4", digits: []int{9, 9, 9}, expected: []int{1, 0, 0, 0}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := PlusOne(tt.digits)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
