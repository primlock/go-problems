/*
https://leetcode.com/problems/single-number/description/

The cheat code is to use the XOR bitwise operator.

Given an array where every element appears twice except for one element, you
can find the non-repeating element by XOR-ing all elements together. This works
because when you XOR all the elements together, pairs of identical elements will
cancel each other out.
*/

package array

func SingleNumber(nums []int) int {
	remainder := 0

	// XOR on every num in nums
	for _, num := range nums {
		remainder ^= num
	}

	return remainder
}
