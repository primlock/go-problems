// https://leetcode.com/problems/decode-string

package stack

import "strings"

func DecodeString(s string) string {
	k := 0
	repStack := make([]int, 0)
	strStack := make([]string, 0)
	currentStr := ""

	// Handle each character based on it's type.
	for _, char := range s {
		if char >= '0' && char <= '9' {
			// Accumulate the number, accounting for multi-digits.
			k = k*10 + int(char-'0')
		} else if char == '[' {
			// Push the current number of reps and current string onto their stacks.
			repStack = append(repStack, k)
			strStack = append(strStack, currentStr)

			// Reset the counter and string for the next level.
			k = 0
			currentStr = ""
		} else if char == ']' {
			// Pop the current number of reps off the stack.
			reps := repStack[len(repStack)-1]
			repStack = repStack[:len(repStack)-1]

			// Pop the current string off the stack.
			lastStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			// Repeat currentStr 'reps' times and add to the last string.
			currentStr = lastStr + strings.Repeat(currentStr, reps)
		} else {
			// Build up the current string.
			currentStr += string(char)
		}
	}

	return currentStr
}
