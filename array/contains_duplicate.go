// https://leetcode.com/problems/contains-duplicate/description/

package leetcode

func ContainsDuplicate(nums []int) bool {
	occurrenceMap := make(map[int]int)

	for _, num := range nums {
		if _, ok := occurrenceMap[num]; ok {
			return true
		}

		occurrenceMap[num] = 0
	}

	return false
}
