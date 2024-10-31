// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length

package slidingwindow

func MaxVowels(s string, k int) int {
	n := len(s) - 1
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}
	vowel_count := 0

	// Initialize the window - check for vowels
	for i := 0; i < k; i++ {
		if _, ok := vowels[string(s[i])]; ok {
			vowel_count++
		}
	}

	max_vowels := vowel_count

	// Slide the window through the remaining characters
	for i := k; i <= n; i++ {
		// If the character being removed from the window was a vowel
		if _, ok := vowels[string(s[i-k])]; ok {
			vowel_count--
		}

		// If the character being added to the window is a vowel
		if _, ok := vowels[string(s[i])]; ok {
			vowel_count++
		}

		if vowel_count > max_vowels {
			max_vowels = vowel_count
		}
	}

	return max_vowels
}
