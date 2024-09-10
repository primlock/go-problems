// https://www.hackerrank.com/challenges/one-week-preparation-kit-mini-max-sum

package basic

import (
	"fmt"
	"slices"
)

func MiniMaxSum(arr []int32) {
	var sum int64 = 0

	slices.Sort(arr)

	for _, num := range arr {
		sum += int64(num)
	}

	min := int64(sum) - int64(arr[4])
	max := int64(sum) - int64(arr[0])

	fmt.Printf("%d %d\n", min, max)
}
