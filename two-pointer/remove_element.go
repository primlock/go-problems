// https://leetcode.com/problems/remove-element/description/

package leetcode

func truncate(nums []int, val int) []int {
	for i, num := range nums {
		if num == val {
			nums[i] = nums[len(nums)-1]
			return nums[:len(nums)-1]
		}
	}
	return nums
}

func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	removals := 0

	for _, num := range nums {
		if num == val {
			removals++
		}
	}

	totalRemovals := len(nums) - removals

	for i := 0; i < removals; i++ {
		nums = truncate(nums, val)
	}

	return totalRemovals
}
