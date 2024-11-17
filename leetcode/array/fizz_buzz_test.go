package array

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected []string
	}{
		{name: "1", n: 3, expected: []string{"1", "2", "Fizz"}},
		{name: "2", n: 5, expected: []string{"1", "2", "Fizz", "4", "Buzz"}},
		{name: "3", n: 15, expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := FizzBuzz(tt.n)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
