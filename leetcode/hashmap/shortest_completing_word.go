// https://leetcode.com/problems/shortest-completing-word

package hashmap

func createFrequencyMap(word string) map[byte]int {
	letters := make(map[byte]int, 0)

	for i := 0; i < len(word); i++ {
		letter := word[i]
		// Ignore numbers and spaces and treat letters as case insensitive.
		if letter > '@' {
			if letter < 'a' {
				letters[letter+32]++
			} else {
				letters[letter]++
			}
		}
	}

	return letters
}

func ShortestCompletingWord(licensePlate string, words []string) string {
	// If a letter appears more than once in licensePlate, then it must appear in the word the same number of times or more.
	// If there are multiple shortest completing words, return the first one that occurs in words.

	shortest := ""

	// Build a list of letters required from license plate.
	letters := createFrequencyMap(licensePlate)

	// For each word in words, verify that the word contains all the letters from license plate.
	for _, word := range words {
		wordMap := createFrequencyMap(word)
		isCompleting := true

		// Compare the required letter map to the word letter map
		for k, v := range letters {
			quantity, ok := wordMap[k]

			// If letter exists, validate the quantity.
			if ok {
				// If there are more of the required letter in word than in license plate, that's ok.
				if quantity < v {
					isCompleting = false
					break
				}
			} else {
				// If the letter doesn't exist, it isn't a completing word.
				isCompleting = false
				break
			}
		}

		// Store the completing word in a variable and replace it when you verify a shorter completing word.
		if isCompleting {
			if shortest == "" {
				shortest = word
			} else {
				if len(word) < len(shortest) {
					shortest = word
				}
			}
		}
	}

	return shortest
}
