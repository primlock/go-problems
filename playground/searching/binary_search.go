package playground

// Time: O(log n), Space: O(1)
func BinarySearch(array []int, low, high, target int) int {
	// While we haven't explored all possibilities
	if low <= high {
		// Find the mid point
		mid := low + (high-low)/2

		if array[mid] == target {
			return mid
		} else if array[mid] > target {
			// Search left portion
			return BinarySearch(array, low, mid-1, target)
		} else {
			// Search right portion
			return BinarySearch(array, mid+1, high, target)
		}
	}

	return -1
}
