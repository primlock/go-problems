/*
	https://www.hackerrank.com/challenges/one-week-preparation-kit-lonely-integer/problem

	Calculate the ABSOLUTE difference between the sums of the square matrix diagonals
	Return the absolute difference
	-------------
	Iterate over each row of the square
		- primary_diagonal starts at index 0 and increases to n-1
		- secondary_diagonal starts at index n-1 and decreases to 0

	Calculate the absolute value with math.Abs()

	Iterating just once:
	i = 0 | i = (n - 1) - i = 2
	i = 1 | i = (n - 1) - i = 1
	i = 2 | i = (n - 1) - i = 0
*/

package basic

import "math"

func DiagonalDifference(arr [][]int32) int32 {
	n := len(arr[0])
	in := 0

	var primary_diagonal int32
	var secondary_diagonal int32

	// Go through each row in the square matrix
	for i := 0; i < n; i++ {
		primary_diagonal += arr[i][in]
		secondary_diagonal += arr[i][(n-1)-in]
		in++
	}

	return int32(math.Abs(float64(primary_diagonal) - float64(secondary_diagonal)))
}
