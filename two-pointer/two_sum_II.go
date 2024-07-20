// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

package leetcode

func TwoSum(numbers []int, target int) []int {
	// Handle base case
	if len(numbers) == 2 {
		return []int{1, 2}
	}

	left := 0
	right := len(numbers) - 1

	for left < right {
		if numbers[left]+numbers[right] > target {
			// ^ Too high: decrement right
			right--
		} else if numbers[left]+numbers[right] < target {
			// ^ Too low: increment left
			left++
		} else {
			// ^ found the target!
			return []int{left + 1, right + 1}
		}
	}

	return nil
}
