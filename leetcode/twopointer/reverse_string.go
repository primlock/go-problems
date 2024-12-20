// https://leetcode.com/problems/reverse-string

package twopointer

func reverseString(s []byte) []byte {
	left, right := 0, len(s)-1

	// Use the swap algorithm to swap the left and right placement of chars
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	return s
}
