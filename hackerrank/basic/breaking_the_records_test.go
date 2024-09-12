package basic

import (
	"reflect"
	"testing"
)

func TestBreakingRecords(t *testing.T) {
	testCases := []struct {
		name     string
		scores   []int32
		expected []int32
	}{
		{name: "1", scores: []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}, expected: []int32{2, 4}},
		{name: "2", scores: []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}, expected: []int32{4, 0}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := BreakingRecords(tt.scores)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
