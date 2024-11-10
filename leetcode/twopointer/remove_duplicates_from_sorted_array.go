package twopointer

func RemoveDuplicates(nums []int) int {
	if len(nums) == 1 {
		return len(nums)
	}

	slow := 1

	for fast := 1; fast < len(nums); fast++ {
		// If this element is unique from the last one
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}
