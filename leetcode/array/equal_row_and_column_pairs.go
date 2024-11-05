// https://leetcode.com/problems/equal-row-and-column-pairs

package array

import "slices"

func EqualPairs(grid [][]int) int {
	pairs := 0
	columns := make([][]int, len(grid))

	// Build a slice of the columns to compare against.
	for _, row := range grid {
		for j, el := range row {
			columns[j] = append(columns[j], el)
		}
	}

	// For every row in the grid, compare it to every column.
	for _, row := range grid {
		for _, col := range columns {
			if slices.Equal(row, col) {
				pairs++
			}
		}
	}

	return pairs
}
