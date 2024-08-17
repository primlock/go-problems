package array

import (
	"reflect"
	"testing"
)

func TestTopKFrequentElements(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{name: "1", nums: []int{1, 1, 1, 2, 2, 3}, k: 2, expected: []int{1, 2}},
		{name: "2", nums: []int{1}, k: 1, expected: []int{1}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := TopKFrequent(tt.nums, tt.k)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
