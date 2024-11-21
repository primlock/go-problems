// https://leetcode.com/problems/find-the-difference/

package hashmap

func FindTheDifference(s string, t string) byte {
	// String t is all the letters in string s with an additional letter.

	// Create a frequency map of all characters in s.
	sMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}

	// Create a frequency map of all charcters in t.
	tMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	// Compare both maps by iterating over t
	for k, v := range tMap {
		// Find the val that doesn't exist or doesn't match the frequency in s.
		val, ok := sMap[k]
		if !ok || val != v {
			return k
		}
	}

	return ' '
}
