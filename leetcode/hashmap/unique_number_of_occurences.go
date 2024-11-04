// https://leetcode.com/problems/unique-number-of-occurrences

package hashmap

func UniqueOccurrences(arr []int) bool {
	// Make a map to record the nums found in arr with their occurences
	nums_map := make(map[int]int, 0)
	for _, num := range arr {
		nums_map[num]++
	}

	// Make a map to record the unique number of occurences
	occur_map := make(map[int]int)
	for key, val := range nums_map {
		// Does the occurence already exist
		_, ok := occur_map[val]
		if ok {
			return false
		}
		occur_map[val] = key
	}

	return true
}
