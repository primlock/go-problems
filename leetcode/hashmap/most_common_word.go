package hashmap

import (
	"slices"
	"strings"
)

func filterString(s string) string {
	// Cast each word to lowercase for comparison.
	lower := strings.ToLower(s)
	result := make([]byte, 0)

	// Do filtering to eliminate the symbols (!?',;.).
	for i := 0; i < len(lower); i++ {
		if lower[i] > '`' && lower[i] < '{' {
			result = append(result, lower[i])
		} else if lower[i] == ' ' {
			result = append(result, lower[i])
		} else if i < len(lower)-1 && lower[i] == ',' && lower[i+1] != ' ' {
			// Annoying edge case where a comma exists but no following space (a,b).
			result = append(result, ' ')
		}
	}

	return string(result)
}

func MostCommonWord(paragraph string, banned []string) string {
	filtered := filterString(paragraph)
	words := strings.Split(filtered, " ")

	// Creating a frequency map of terms (in lowercase) from paragraph.
	paragraphFrequencyMap := map[string]int{}
	for _, word := range words {
		paragraphFrequencyMap[word]++
	}

	// Find the most frequent word that is not banned.
	mostFrequentWord := ""
	max := 0
	for k, v := range paragraphFrequencyMap {
		// Compare this word to the list of banned words
		if slices.Contains(banned, k) {
			continue
		}

		// Record the most frequent
		if v > max {
			mostFrequentWord = k
			max = v
		}
	}

	return mostFrequentWord
}
