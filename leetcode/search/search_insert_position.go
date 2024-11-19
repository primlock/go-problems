// https://leetcode.com/problems/search-insert-position

package search

func binarySearch(array []int, start, end, target int) int {
	if start > end {
		return start
	}

	mid := (start + end) / 2

	if array[mid] == target {
		return mid
	} else if array[mid] > target {
		return binarySearch(array, start, mid-1, target)
	} else {
		return binarySearch(array, mid+1, end, target)
	}
}

func SearchInsert(nums []int, target int) int {
	// Binary Search runs O(log n)
	return binarySearch(nums, 0, len(nums)-1, target)
}
