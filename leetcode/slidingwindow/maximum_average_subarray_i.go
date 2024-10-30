// https://leetcode.com/problems/maximum-average-subarray-i

package slidingwindow

func FindMaxAverage(nums []int, k int) float64 {
	n := len(nums)

	// Find the starting average of the subarray
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	average := float64(windowSum) / float64(k)

	// Use the first average as the max
	maxAverage := average

	// Continue sliding the window to look for a new max
	for i := k; i < n; i++ {
		// Add the next element and remove the first element in the current window
		windowSum += nums[i] - nums[i-k]

		// Update max if average is greater
		average = float64(windowSum) / float64(k)
		if average > maxAverage {
			maxAverage = average
		}
	}

	return maxAverage
}
