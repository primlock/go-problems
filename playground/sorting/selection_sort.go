package playground

// Time: O(n²), Space: O(1)
func SelectionSort(array []int, n int) []int {
	// Iterate through all elements
	for i := 0; i < (n - 1); i++ {
		// Min index is assumed to be the current iteration
		min_i := i

		// Look ahead at the remaining elements
		for j := (i + 1); j < n; j++ {
			// Determine if we have a new min index by comparing elements
			if array[j] < array[min_i] {
				min_i = j
			}
		}

		// Swap the elements
		if min_i != i {
			array = Swap(array, i, min_i)
		}
	}

	return array
}
