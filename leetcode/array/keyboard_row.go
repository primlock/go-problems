// https://leetcode.com/problems/keyboard-row/

package array

func verifyLowercase(b byte) byte {
	wordVal := int(b)

	if wordVal < 97 {
		// Its lowercase, so we need to add 32 to get it's lowercase counterpart
		wordVal += 32
	}

	return byte(wordVal)
}

func FindWords(words []string) []string {
	// The first letter of the word determine which row it's committed to.
	keyboardMap := map[byte]int{
		'q': 1,
		'w': 1,
		'e': 1,
		'r': 1,
		't': 1,
		'y': 1,
		'u': 1,
		'i': 1,
		'o': 1,
		'p': 1,
		'a': 2,
		's': 2,
		'd': 2,
		'f': 2,
		'g': 2,
		'h': 2,
		'j': 2,
		'k': 2,
		'l': 2,
		'z': 3,
		'x': 3,
		'c': 3,
		'v': 3,
		'b': 3,
		'n': 3,
		'm': 3,
	}

	sameRowWords := make([]string, 0)

	for _, word := range words {
		committedRow := keyboardMap[verifyLowercase(word[0])]
		isSameRow := true

		for i := 0; i < len(word); i++ {
			// If the letter is not in commited row, move to the next word
			val := keyboardMap[verifyLowercase(word[i])]
			if val != committedRow {
				isSameRow = false
				break
			}
		}

		if isSameRow {
			sameRowWords = append(sameRowWords, word)
		}
	}

	return sameRowWords
}
