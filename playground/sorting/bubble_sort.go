package playground

// Time: O(n²), Space: O(1)
func BubbleSort(array []int, n int) []int {
	swapped := false

	for i := 0; i < (n - 1); i++ {
		swapped = false

		// Iterate up to the last sorted element
		for j := 0; j < (n-i)-1; j++ {
			// If the left side is greater than the right side
			if array[j] > array[j+1] {
				array = Swap(array, j, j+1)
				swapped = true
			}
		}

		// Break away early if nothing needs to be swapped
		if !swapped {
			break
		}
	}

	return array
}
