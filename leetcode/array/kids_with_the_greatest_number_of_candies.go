// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies

package array

import "math"

func KidsWithCandies(candies []int, extraCandies int) []bool {
	// Find the max number of candies
	max := math.MinInt
	for _, num := range candies {
		if num > max {
			max = num
		}
	}

	// Build an array of bool by adding candy to the max and evaluating
	result := make([]bool, 0)
	for _, candy := range candies {
		if candy+extraCandies >= max {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}
