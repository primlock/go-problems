// https://leetcode.com/problems/string-compression

package array

import (
	"fmt"
)

func Compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}

	s := make([]byte, 0)
	streak := 1

	for i, char := range chars {
		// If this char is the same as the last char, increase the streak
		if i > 0 && char == chars[i-1] {
			streak++
		}

		// If the streak ends, append its compressed bytes to s
		if i > 0 && char != chars[i-1] {
			s = append(s, chars[i-1])

			if streak > 1 {
				nums := fmt.Sprintf("%d", streak)

				for _, num := range nums {
					s = append(s, byte(num))
				}
			}

			streak = 1
		}

		// If this is the last char
		if i == len(chars)-1 {
			s = append(s, char)

			if streak > 1 {
				nums := fmt.Sprintf("%d", streak)

				for _, num := range nums {
					s = append(s, byte(num))
				}
			}
		}

	}

	// Store s in chars
	copy(chars, s)
	chars = chars[:len(s)]

	return len(chars)
}
