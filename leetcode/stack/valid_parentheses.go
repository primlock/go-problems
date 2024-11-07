// https://leetcode.com/problems/valid-parentheses

package stack

func IsValid(s string) bool {
	stack := []byte{}
	closing := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}

	for _, char := range s {
		// If this is an opening char, push onto the stack
		if char == '[' || char == '{' || char == '(' {
			stack = append(stack, byte(char))
			continue
		}

		// If this is a closing char, pop its pair off the stack
		if len(stack) > 0 && stack[len(stack)-1] == closing[byte(char)] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
