// https://leetcode.com/problems/determine-if-two-strings-are-close

package hashmap

import (
	"maps"
	"slices"
)

func operationOne(word1 string, word2 string) bool {
	word1_map := make(map[string]int, 0)
	word2_map := make(map[string]int, 0)

	for i := 0; i < len(word1); i++ {
		word1_map[string(word1[i])]++
		word2_map[string(word2[i])]++
	}

	return maps.Equal(word1_map, word2_map)
}

func operationTwo(word1 string, word2 string) bool {
	word1_map := make(map[string]int)
	word2_map := make(map[string]int)

	// Create frequency maps for each word
	for _, char := range word1 {
		word1_map[string(char)]++
	}
	for _, char := range word2 {
		word2_map[string(char)]++
	}

	// Make sure both words have the same unique characters
	if len(word1_map) != len(word2_map) {
		return false
	}
	for char := range word1_map {
		_, ok := word2_map[char]
		if !ok {
			return false
		}
	}

	// Put both sets of frequencies in a slice and sort them
	word1_freq := make([]int, 0, len(word1_map))
	word2_freq := make([]int, 0, len(word2_map))

	for _, val := range word1_map {
		word1_freq = append(word1_freq, val)
	}

	for _, val := range word2_map {
		word2_freq = append(word2_freq, val)
	}

	slices.Sort(word1_freq)
	slices.Sort(word2_freq)

	// Compare for equality
	return slices.Equal(word1_freq, word2_freq)
}

func CloseStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	close := operationOne(word1, word2)
	if close {
		return true
	}

	return operationTwo(word1, word2)
}
