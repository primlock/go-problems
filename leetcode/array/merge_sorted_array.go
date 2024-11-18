// https://leetcode.com/problems/merge-sorted-array/

package array

import "slices"

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	// Add all the elements into nums1
	j := 0
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[j]
		j++
	}

	// Call sort to arrange in non-decreasing
	slices.Sort(nums1)

	return nums1
}
