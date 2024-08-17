// https://leetcode.com/problems/longest-consecutive-sequence/description/

package array

func LongestConsecutiveSequence(nums []int) int {
	maxStreak := 0

	// Build a set of unique values
	uniqueMap := make(map[int]bool)

	for _, num := range nums {
		if _, ok := uniqueMap[num]; !ok {
			uniqueMap[num] = true
		}
	}

	// Find the start of a sequence
	for key := range uniqueMap {
		if _, ok := uniqueMap[key-1]; ok {
			continue
		}

		// Continue counting the streak until num is not in sequence
		streak := 0
		for {
			if _, ok := uniqueMap[key+streak]; ok {
				streak++
			} else {
				maxStreak = max(streak, maxStreak)
				break
			}
		}
	}

	return maxStreak
}
