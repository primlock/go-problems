// https://leetcode.com/problems/container-with-most-water

package twopointer

import "math"

func MaxArea(height []int) int {
	// Start at opposite sides of the container
	left := 0
	right := len(height) - 1
	max := 0

	for left < right {
		// Can't slant, so we need to take the min height
		min_h := math.Min(float64(height[left]), float64(height[right]))

		w := right - left

		// Determine how much water we can contain
		area := int(min_h) * w

		if area > max {
			max = area
		}

		// Move the smaller side or left if they're the same
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return max
}
