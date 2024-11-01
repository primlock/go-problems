package prefixsum

func LargestAltitude(gain []int) int {
	// Create a prefix sum for the gain
	prefix := make([]int, len(gain)+1)
	for i := 0; i < len(gain); i++ {
		prefix[i+1] = prefix[i] + gain[i]
	}

	// Find the max gain by iterating over prefix and storing max
	max := 0
	for i := 0; i < len(prefix); i++ {
		if prefix[i] > max {
			max = prefix[i]
		}
	}

	return max
}
