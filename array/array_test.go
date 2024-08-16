package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{name: "1", nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
		{name: "2", nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
		{name: "3", nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{name: "1", nums: []int{1, 2, 3, 1}, expected: true},
		{name: "2", nums: []int{1, 2, 3, 4}, expected: false},
		{name: "3", nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, expected: true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsDuplicate(tt.nums)
			want := tt.expected

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func TestValidAnagram(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{name: "1", s: "anagram", t: "nagaram", expected: true},
		{name: "2", s: "rat", t: "car", expected: false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidAnagram(tt.s, tt.t)
			want := tt.expected

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

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

func TestProductOfArrayExceptSelf(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{name: "1", nums: []int{1, 2, 3, 4}, expected: []int{24, 12, 8, 6}},
		{name: "2", nums: []int{-1, 1, 0, -3, 3}, expected: []int{0, 0, 9, 0, 0}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := ProductOfArrayExceptSelf(tt.nums)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func TestLongestConsecutiveSequence(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{name: "1", nums: []int{100, 4, 200, 1, 3, 2}, expected: 4},
		{name: "2", nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, expected: 9},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestConsecutiveSequence(tt.nums)
			want := tt.expected

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}

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
