// https://leetcode.com/problems/length-of-last-word

package lstring

import "strings"

func lengthOfLastWord(s string) int {
	// Filter down the string if it contains spaces.
	words := strings.Fields(s)

	// Get the length of that last word in the slice.
	return len(words[len(words)-1])
}
