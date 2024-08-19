/*
https://leetcode.com/problems/plus-one/description/

Keep looping over the array while we still have a remainder.

If the element we're at is less than 9, increment that element and return
the result.

If element is 9, set that element to 0 and increment the
carry for another loop. If we reach the first element (index 0) and
still have a carry, set the current element to 0 and append a 1 onto
the front of the array.
*/

package array

func PlusOne(digits []int) []int {
	carry := 1
	i := len(digits) - 1

	for carry > 0 {
		carry--

		if digits[i] == 9 {
			// Roll the 9 over to a 0
			digits[i] = 0
			carry++

			// Most significan number
			if i == 0 {
				// Prepend a 1 if we are on the MSN
				digits = append([]int{1}, digits...)
				carry--
			}
		} else {
			digits[i]++
		}

		if i > 0 {
			i--
		}
	}

	return digits
}
