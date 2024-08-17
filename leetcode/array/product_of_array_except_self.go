// https://leetcode.com/problems/product-of-array-except-self/description/

package array

func ProductOfArrayExceptSelf(nums []int) []int {
	// Get the values for the prefix array
	prefix := make([]int, len(nums))
	for i, num := range nums {
		// guard this
		if i == 0 {
			prefix[i] = 1 * num
		} else {
			prefix[i] = prefix[i-1] * num
		}
	}

	// Get the values for the postfix array
	postfix := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			postfix[i] = 1 * nums[i]
		} else {
			postfix[i] = postfix[i+1] * nums[i]
		}
	}

	// Combine the prefix array with the postfix array
	result := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			result[i] = 1 * postfix[i+1]
		} else if i == len(nums)-1 {
			result[i] = prefix[i-1] * 1
		} else {
			result[i] = prefix[i-1] * postfix[i+1]
		}
	}

	return result
}
