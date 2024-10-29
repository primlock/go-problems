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
	first := 0
	second := 0

	for second != len(nums)-1 {
		if first < len(nums) && nums[first] == 0 {
			for second != len(nums)-1 && nums[second] == 0 {
				second++
			}

			nums[first], nums[second] = nums[second], nums[first]
			first++
		} else {
			first++
			second++
		}
	}

	return nums
}
