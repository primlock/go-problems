// https://leetcode.com/problems/guess-number-higher-or-lower

package search

// The solution has been modified slightly from the submission to provide guess with the picked number.

func guess(num, pick int) int {
	if num == pick {
		return 0
	} else if num > pick {
		return -1
	} else {
		return 1
	}
}

func GuessNumber(n, pick int) int {
	low := 1
	high := n

	for low <= high {
		g := low + (high-low)/2
		res := guess(g, pick)

		if res == 0 {
			return g
		} else if res == -1 {
			high = g - 1
		} else {
			low = g + 1
		}
	}

	return -1
}
