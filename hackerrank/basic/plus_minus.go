// https://www.hackerrank.com/challenges/one-week-preparation-kit-plus-minus/problem

package basic

import "fmt"

func PlusMinus(arr []int32) {
	pos := 0
	neg := 0
	zero := 0
	n := len(arr)

	for _, num := range arr {
		if num > 0 {
			pos += 1
		} else if num < 0 {
			neg += 1
		} else {
			zero += 1
		}
	}

	// float64 cast must be done before division
	fmt.Printf("%.6f\n", float64(pos)/float64(n))
	fmt.Printf("%.6f\n", float64(neg)/float64(n))
	fmt.Printf("%.6f\n", float64(zero)/float64(n))
}
