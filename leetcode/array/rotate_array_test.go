package array

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{name: "1", nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, expected: []int{5, 6, 7, 1, 2, 3, 4}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Rotate(tt.nums, tt.k)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
