// https://leetcode.com/problems/valid-palindrome/description/

package leetcode

import (
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile("[^a-z0-9]+")

func ValidPalindrome(s string) bool {
	// Clean our string
	ls := strings.ToLower(s)
	filtered := nonAlphanumericRegex.ReplaceAllString(ls, "")

	if len(filtered) == 0 {
		return true
	}

	left := 0
	right := len(filtered) - 1

	// Use two pointer to interate over filtered
	for left < right {
		// If left doesn't match the right, not a palindrome
		if filtered[left] != filtered[right] {
			return false
		}

		left++
		right--
	}

	return true
}
