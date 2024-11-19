package array

import (
	"reflect"
	"testing"
)

func TestShortestToChar(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		c        byte
		expected []int
	}{
		{name: "1", s: "loveleetcode", c: 'e', expected: []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}},
		{name: "2", s: "aaab", c: 'b', expected: []int{3, 2, 1, 0}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := ShortestToChar(tt.s, tt.c)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
