// https://leetcode.com/problems/successful-pairs-of-spells-and-potions

package search

import "slices"

func SuccessfulPairs(spells []int, potions []int, success int64) []int {
	// successful = spells[i] * potions[j] >= success
	// pairs[i] = # potions that will form a successful pair with spell[i]

	pairs := make([]int, len(spells))

	// Sort the potions array
	slices.Sort(potions)

	spell := 0
	for spell < len(spells) {
		// Use binary search to determine successful potions
		low, high := 0, len(potions)-1

		for low <= high {
			mid := low + (high-low)/2

			if spells[spell]*potions[mid] < int(success) {
				// If the current position of the pointers is too low, bring up low
				low = mid + 1
			} else {
				// Successful or too high, bring down high
				high = mid - 1
			}
		}

		// Number of successful pairs is equal to the total number of potions minus the minimum to be atleast success
		pairs[spell] = len(potions) - low
		spell++
	}

	return pairs
}
