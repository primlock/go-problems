// https://leetcode.com/problems/delete-characters-to-make-fancy-string

package array

import "strings"

func MakeFancyString(s string) string {
	count := 1
	result := []string{}

	// Append the character occurences that are less than 3 to a string slice
	for i, char := range s {
		if i == 0 {
			result = append(result, string(char))
			continue
		}

		// If this character is the same as the last
		if char == rune(s[i-1]) {
			count++
			if count < 3 {
				result = append(result, string(char))
			}
		} else {
			count = 1
			result = append(result, string(char))
		}
	}

	return strings.Join(result, "")
}
