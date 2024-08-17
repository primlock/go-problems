package array

func Rotate(nums []int, k int) []int {
	n := len(nums)
	rotated := []int{}

	// find the minimum number of rotations
	k = k % n

	// translate the left rotatios into right rotations
	rotations := n - k

	// add from k to the end into the new array
	for i := rotations; i < n; i++ {
		rotated = append(rotated, nums[i])
	}

	// add from 0 to k in the new array
	for i := 0; i < rotations; i++ {
		rotated = append(rotated, nums[i])
	}

	return rotated
}
