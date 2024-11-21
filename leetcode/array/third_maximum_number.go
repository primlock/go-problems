// https://leetcode.com/problems/third-maximum-number

package array

import "slices"

func ThirdMax(nums []int) int {
	// Return the THIRD DISTINCT maximum number, if it doesn't exist return the maximum number

	slices.Sort(nums)

	// Use an array to keep track of encountered numbers
	encountered := make([]int, 0)

	// Iterate over nums, adding the distinct nums into encountered
	for _, num := range nums {
		if !slices.Contains(encountered, num) {
			encountered = append(encountered, num)
		}
	}

	// Return the third maximum or the maximum if it doesn't exist.
	if len(encountered) >= 3 {
		return encountered[len(encountered)-3]
	} else {
		return encountered[len(encountered)-1]
	}
}
