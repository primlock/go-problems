package array

import "strings"

func MakeFancyString(s string) string {
	count := 1
	result := []string{}

	for i, char := range s {
		if i == 0 {
			result = append(result, string(char))
			continue
		}

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
