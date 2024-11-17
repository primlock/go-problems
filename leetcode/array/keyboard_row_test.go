package array

import (
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	testCases := []struct {
		name     string
		words    []string
		expected []string
	}{
		{name: "1", words: []string{"Hello", "Alaska", "Dad", "Peace"}, expected: []string{"Alaska", "Dad"}},
		{name: "2", words: []string{"omk"}, expected: []string{}},
		{name: "3", words: []string{"adsdf", "sfd"}, expected: []string{"adsdf", "sfd"}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := FindWords(tt.words)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
