package prefixsum

func PivotIndex(nums []int) int {
	// Create a prefix sum
	prefix := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	// Iterate over indices of nums
	for i := 0; i < len(nums); i++ {
		// Sum of elements to the left of index i
		left := prefix[i]
		// Sum of elements to the right of index i
		right := prefix[len(nums)] - prefix[i+1]

		if left == right {
			return i
		}
	}

	return -1
}
