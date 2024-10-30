// https://leetcode.com/problems/max-number-of-k-sum-pairs

package twopointer

import "slices"

func MaxOperations(nums []int, k int) int {
	left := 0
	right := len(nums) - 1
	operations := 0

	// Sort the array so we know which side to adjust
	slices.Sort(nums)

	for left < right {
		if nums[left]+nums[right] > k {
			// Too high, reduce right
			right--
		} else if nums[left]+nums[right] < k {
			// Too low, increase left
			left++
		} else {
			// K-sum pair found
			operations++
			// Simulate removal by adjusting left and right
			left++
			right--
		}
	}

	return operations
}
