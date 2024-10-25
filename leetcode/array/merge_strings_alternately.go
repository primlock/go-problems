// https://leetcode.com/problems/merge-strings-alternately/

package array

func MergeAlternately(word1 string, word2 string) string {
	result := ""

	// Iterate over the shorter string
	shorter := min(len(word1), len(word2))

	for i := 0; i < shorter; i++ {
		result += string(word1[i])
		result += string(word2[i])
	}

	// Append the remaining characters
	result += word1[shorter:]
	result += word2[shorter:]

	return result
}
