// https://leetcode.com/problems/max-consecutive-ones-iii

package slidingwindow

func LongestOnes(nums []int, k int) int {
	left, right := 0, 0
	max_ones := 0
	zeroes := 0

	// Slide the window over the entire array
	for right < len(nums) {
		// Expand the window if we have zeroes to flip
		if nums[right] == 0 {
			zeroes++
		}
		right++

		// Shrink the window if we have reached our flip limit
		for zeroes > k {
			if nums[left] == 0 {
				zeroes--
			}
			left++
		}

		// Check to see if we have a new max
		ones := right - left
		if ones > max_ones {
			max_ones = ones
		}
	}

	return max_ones
}
