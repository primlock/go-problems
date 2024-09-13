/*
https://leetcode.com/problems/move-zeroes/

Move the zeros to the end of the array without changing the 'relative' order of the
non zero elements.
You need to do this in place without making a copy of the array.
---------------------------------------------
Start a left and a right pointer out at zero.
Loop while right pointer is not at the end of the array.
If left pointer is at a zero, advance right pointer until its at a non-zero
Swap the value at left pointer with the value at right pointer
*/

package twopointer

// Returning []int for the purpose of evaluation
func MoveZeroes(nums []int) []int {
	left := 0
	right := 0

	for right != (len(nums) - 1) {
		if left < len(nums) && nums[left] == 0 {
			for right != (len(nums)-1) && nums[right] == 0 {
				right++
			}

			tmp := nums[left]
			nums[left] = nums[right]
			nums[right] = tmp

			left++
		} else {
			left++
			right++
		}
	}

	return nums
}
