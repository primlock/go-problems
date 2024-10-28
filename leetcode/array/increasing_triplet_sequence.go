// https://leetcode.com/problems/increasing-triplet-subsequence

package array

import "math"

func IncreasingTriplet(nums []int) bool {
	// Track the first and second values of our triplet
	first := math.MaxInt
	second := math.MaxInt

	/*
		First is updated if num is less than equal to first.
		Second is updated if num is greater than first but less than equal to second.
		Triplet found if num is greater than first and greater than second.
	*/

	for _, num := range nums {
		if num <= first {
			first = num
		} else if num > first && num <= second {
			second = num
		} else if num > first && num > second {
			return true
		}
	}

	return false
}
