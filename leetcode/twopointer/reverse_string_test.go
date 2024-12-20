package twopointer

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		name     string
		s        []byte
		expected []byte
	}{
		{name: "1", s: []byte("hello"), expected: []byte("olleh")},
		{name: "2", s: []byte("Hannah"), expected: []byte("hannaH")},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseString(tt.s)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
