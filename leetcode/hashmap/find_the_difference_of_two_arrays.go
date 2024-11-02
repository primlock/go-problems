package hashmap

import "slices"

func FindDifference(nums1 []int, nums2 []int) [][]int {
	num_map := make(map[int][]int)
	answers := make([][]int, 2)

	// Add all the nums found in nums1
	num_map[1] = append(num_map[1], nums1...)

	// Add all the nums found in nums2
	num_map[2] = append(num_map[2], nums2...)

	for _, num := range num_map[1] {
		// If it's not in nums2
		if !slices.Contains(num_map[2], num) {
			if !slices.Contains(answers[0], num) {
				answers[0] = append(answers[0], num)
			}
		}
	}

	for _, num := range num_map[2] {
		// If it's not in nums1
		if !slices.Contains(num_map[1], num) {
			if !slices.Contains(answers[1], num) {
				answers[1] = append(answers[1], num)
			}
		}
	}

	return answers
}
