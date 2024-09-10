/*
https://www.hackerrank.com/challenges/one-week-preparation-kit-lonely-integer/problem

We need to find the unique integer in the array. An easy way to do that is to use
XOR on each element until we're left with only the unique. This is bit manipulation
which is explained in bit-mainipulation.md.

*/

package basic

func LonelyInteger(a []int32) int32 {
	var unique int32

	for _, el := range a {
		unique ^= el
	}

	return unique
}
