// https://leetcode.com/problems/circular-sentence

package array

import "strings"

func IsCircularSentence(sentence string) bool {
	s := strings.Fields(sentence)

	// Word length of 1, compare the first and last character of the only word.
	if len(s) == 1 {
		return s[0][0] == s[0][len(s[0])-1]
	}

	// Check that the first character of the first word equals the last character in the last word.
	if s[0][0] != s[len(s)-1][len(s[len(s)-1])-1] {
		return false
	}

	// Check that the last character of the current word equals the first character of the next word.
	for i, word := range s {
		if i < len(s)-1 {
			last_char_current_word := word[len(word)-1]
			first_char_next_word := s[i+1][0]

			if last_char_current_word != first_char_next_word {
				return false
			}
		}
	}

	return true
}
