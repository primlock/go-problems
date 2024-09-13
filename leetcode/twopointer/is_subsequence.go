/*
https://leetcode.com/problems/is-subsequence/description/

Determine if the string 's' is a subsequence of the string 't'.
Relative order must be maintained.
Return true if 's' is a subsequence, false otherwise.
--------------------------------------------
Place a left pointer at the first index of 's' and a right pointer at the first index of 't'.

Create a loop that runs while we still have characters to iterate in 't'.

If left pointer character is equal to right pointer character, advance left pointer.
- Otherwise, advance right pointer only.

If left pointer is equal to the length of s, we found the substring.
If right pointer is equal to the length of t before left pointer reaches the end, we didn't find a substring.
*/

package twopointer

func IsSubsequence(s string, t string) bool {
	// Handle the edges
	if len(s) == 0 && len(t) == 0 {
		return true
	} else if len(s) == 0 && len(t) != 0 {
		return true
	} else if len(s) != 0 && len(t) == 0 {
		return false
	}

	left := 0
	right := 0

	for right < len(t) {
		if s[left] == t[right] {
			left++

			if left == len(s) {
				return true
			}
		}
		right++
	}

	return false
}
