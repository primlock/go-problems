package hashmap

func canConstruct(ransomNote string, magazine string) bool {
	// Is the enough letters in magazine to make the ransom note?

	// Build the hashmap to take inventory of what we have.
	m := map[byte]int{}
	for i := 0; i < len(magazine); i++ {
		m[magazine[i]]++
	}

	// Determine if we can make the ransom note.
	for i := 0; i < len(ransomNote); i++ {
		count, ok := m[ransomNote[i]]
		if !ok {
			// The letter isn't in the magazine
			return false
		} else if count == 0 {
			// We don't have enough letters left
			return false
		}

		m[ransomNote[i]]--
	}

	return true
}
