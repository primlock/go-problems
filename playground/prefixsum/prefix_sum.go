/*
Prefix Sum Technique

A prefix sum array is a way to keep cumulative totals up to each position in an array, making it easier to answer queries about sums of subarrays. For any position i, the prefix sum at that point is the sum of all elements from the beginning of the array up to i. Once you have the prefix sum, you can quickly calculate the sum of any subarray by subtracting prefix sums, which helps avoid recalculating totals repeatedly.

Suppose you have an array of integers, and you want to find the sum of elements between two indices i and j multiple times. Calculating the sum directly with a loop would be inefficient for repeated queries. With a prefix sum array, you can quickly get any subarray sum without re-summing from scratch.
*/

package playground

func PrefixSum(nums []int, i, j int) int {
	// Create a prefix sum array with an initial 0 for convenience
	prefix := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	// Find the sum in the range i, j efficiently
	return prefix[j+1] - prefix[i]
}
