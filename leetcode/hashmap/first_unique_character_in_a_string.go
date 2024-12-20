package hashmap

func firstUniqueChar(s string) int {
	// Create a frequency map of all the characters
	freq := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	// Find only the uniques from our frequency map
	uniques := make(map[byte]int, 0)
	for k, v := range freq {
		if v == 1 {
			uniques[k]++
		}
	}

	// If there are no uniques, return -1
	if len(uniques) == 0 {
		return -1
	}

	// Go through s to find the first unique and return it's index
	for i := 0; i < len(s); i++ {
		_, ok := uniques[s[i]]
		if ok {
			return i
		}
	}

	return -1
}
