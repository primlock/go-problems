package array

func MissingNumber(nums []int) int {
	sumOfNums := 0
	sumOfRange := 0
	n := len(nums)

	for _, num := range nums {
		sumOfNums += num
	}

	sumOfRange = n * (n + 1) / 2

	return sumOfRange - sumOfNums
}
