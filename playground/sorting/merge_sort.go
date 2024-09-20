package playground

func merge(a, b []int) []int {
	final := make([]int, 0)
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		// Append the element according to desired sort order
		if a[i] < b[i] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	// Append the remaining elements from both slices
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}

// Time: O(n*log(n)), Space: O(n)
func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	// Divide the array in half until you can't anymore
	first := MergeSort(array[:len(array)/2])
	second := MergeSort(array[len(array)/2:])
	return merge(first, second)
}
