// https://leetcode.com/problems/reverse-words-in-a-string

package array

import (
	"slices"
	"strings"
)

func ReverseWords(s string) string {
	// Break it apart by whitespace
	words := strings.Fields(s)

	// Reverse the slice
	slices.Reverse(words)

	// Join the slice back together
	return strings.Join(words, " ")
}
