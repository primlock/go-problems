// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element

package slidingwindow

func LongestSubarray(nums []int) int {
	left, right := 0, 0
	zeroes := 0
	max := 0

	for right < len(nums) {
		// Expand the window if we encounter a 0
		if nums[right] == 0 {
			zeroes++
		}

		right++

		// Shrink the window until we've used only one deletion
		for zeroes > 1 {
			if nums[left] == 0 {
				zeroes--
			}
			left++
		}

		// Determine the size of the subarray of 1's
		length := right - left
		if length > max {
			// length - 1 to account for the removed element
			max = length - 1
		}
	}
	return max
}
