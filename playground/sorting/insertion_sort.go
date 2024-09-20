package playground

// Time: O(n²), Space: O(1)
func InsertionSort(array []int, n int) []int {
	// Start comparing from index 1
	for i := 1; i < n; i++ {
		// Store the value we're comparing
		key := array[i]

		// Use j to iterate backwards from i
		j := i - 1

		// Look backwards from i for elements out of order
		for j >= 0 && array[j] > key {
			// Insert the element in the correct order
			array[j+1] = array[j]
			j--
		}

		array[j+1] = key
	}

	return array
}
