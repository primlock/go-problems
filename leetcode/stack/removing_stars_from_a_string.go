// https://leetcode.com/problems/removing-stars-from-a-string

package stack

func RemoveStars(s string) string {
	stack := make([]byte, 0)

	// Cast the string to a slice of bytes to speed up the code.
	for _, char := range []byte(s) {
		// If we encounter a star, pop the last element from the stack.
		if char == '*' {
			stack = stack[:len(stack)-1]
		} else {
			// Not a star, just append the character.
			stack = append(stack, char)
		}
	}

	return string(stack)
}
