// https://leetcode.com/problems/word-pattern/

package lstring

import (
	"slices"
	"strings"
)

func WordPattern(pattern string, s string) bool {
	// Split the pattern and s into slices
	words := strings.Split(s, " ")
	letters := strings.Split(pattern, "")

	if len(words) != len(letters) {
		return false
	}

	// Find the unique words and letters
	unique_s := make([]string, 0)
	for _, word := range words {
		if !slices.Contains(unique_s, word) {
			unique_s = append(unique_s, word)
		}
	}

	unique_p := make([]string, 0)
	for _, p := range letters {
		if !slices.Contains(unique_p, p) {
			unique_p = append(unique_p, p)
		}
	}

	// Map each unique letter in patter to a unique word in words
	m := map[string]string{}
	n := len(unique_s)
	if len(unique_p) < len(unique_s) {
		n = len(unique_p)
	}

	for i := 0; i < n; i++ {
		m[unique_p[i]] = unique_s[i]
	}

	// Determine if the word in words matches against the letter we assigned it
	for i, word := range words {
		if word != m[letters[i]] {
			return false
		}
	}

	return true
}
