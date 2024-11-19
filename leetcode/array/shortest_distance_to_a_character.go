package array

import "math"

func ShortestToChar(s string, c byte) []int {
	answer := make([]int, len(s))
	occurence := make([]int, 0)

	// Create a list of the indexes of c in s.
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			occurence = append(occurence, i)
		}
	}

	// As we iterate over s, append the minimum value between i and occurence[c]
	for i := 0; i < len(s); i++ {
		min := len(s)
		// For every character in s, find the minimum distance to c.
		for j := 0; j < len(occurence); j++ {
			distance := int(math.Abs(float64(i - occurence[j])))

			if distance < min {
				min = distance
			}
		}

		answer[i] = min
	}

	return answer
}
