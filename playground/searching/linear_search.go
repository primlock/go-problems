package playground

// Time: O(n), Space: O(1)
func LinearSearch(array []int, n, target int) int {
	// Go through the entire list until the element is found
	for i := 0; i < n; i++ {
		if array[i] == target {
			return i
		}
	}

	return -1
}
