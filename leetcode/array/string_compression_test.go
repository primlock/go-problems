package array

import (
	"testing"
)

func TestCompress(t *testing.T) {
	testCases := []struct {
		name     string
		chars    []byte
		expected int
	}{
		{name: "1", chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, expected: 6},
		{name: "2", chars: []byte{'a'}, expected: 1},
		{name: "3", chars: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, expected: 4},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Compress(tt.chars)
			want := tt.expected

			if got != want {
				t.Errorf("size: got %v, want %v", got, want)
			}
		})
	}
}
