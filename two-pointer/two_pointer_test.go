package leetcode

import (
	"reflect"
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected bool
	}{
		{name: "1", s: "A man, a plan, a canal: Panama", expected: true},
		{name: "2", s: "race a car", expected: false},
		{name: "3", s: " ", expected: true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidPalindrome(tt.s)
			want := tt.expected

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name     string
		numbers  []int
		target   int
		expected []int
	}{
		{name: "1", numbers: []int{2, 7, 11, 15}, target: 9, expected: []int{1, 2}},
		{name: "2", numbers: []int{2, 3, 4}, target: 6, expected: []int{1, 3}},
		{name: "3", numbers: []int{-1, 0}, target: -1, expected: []int{1, 2}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.numbers, tt.target)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
