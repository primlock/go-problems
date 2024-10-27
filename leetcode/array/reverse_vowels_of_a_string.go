package array

func ReverseVowels(s string) string {
	vowels := map[string]bool{
		"A": true,
		"E": true,
		"I": true,
		"O": true,
		"U": true,
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}

	// Get a count of vowels to init our array with
	vowel_count := 0
	for i := 0; i < len(s); i++ {
		if _, ok := vowels[string(s[i])]; ok {
			vowel_count++
		}
	}

	// Build the vowel array
	v := make([]rune, vowel_count)
	n := vowel_count
	for _, char := range s {
		if _, ok := vowels[string(char)]; ok {
			v[n-1] = char
			n--
		}
	}

	// Replace the characters
	j := 0
	runes := []rune(s)
	for i, char := range runes {
		if _, ok := vowels[string(char)]; ok {
			runes[i] = v[j]
			j++
		}
	}

	return string(runes)
}
