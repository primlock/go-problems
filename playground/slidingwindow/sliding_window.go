// This program performs a basic sliding window function to find a maximum sum in a subset of values.

package playground

func MaxSumInSubset(nums []int, k int) int {
	n := len(nums)

	// Guard the slice access
	if k > n || k <= 0 {
		return 0
	}

	// Initialize the window sum by counting the first k elements.
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	max := sum

	// Perform a sliding window to check for a new max - starting at k
	for i := k; i < n; i++ {
		// Add the next element into the window and subtract the element leaving the window
		sum += nums[i] - nums[i-k]

		if sum > max {
			max = sum
		}
	}

	return max
}
