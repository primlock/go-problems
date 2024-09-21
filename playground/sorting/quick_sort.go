package playground

// Time: O(n*log(n)), Space: O(1)
func QuickSort(array []int) []int {
	return quickSort(array, 0, len(array)-1)
}

// Divide the array into smaller sub arrays and sort
func quickSort(array []int, low, high int) []int {
	if low < high {
		var p int
		array, p = partition(array, low, high)
		array = quickSort(array, low, p-1)
		array = quickSort(array, p+1, high)
	}
	return array
}

// Partition the array around a pivot element
func partition(array []int, low, high int) ([]int, int) {
	pivot := array[high]
	i := low

	for j := low; j < high; j++ {
		// Swap the elements if they're smaller than the pivot
		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}
	// Place element in it's correct position in the array
	array[i], array[high] = array[high], array[i]
	return array, i
}
