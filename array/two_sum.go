// https://leetcode.com/problems/two-sum/description/

package leetcode

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		diff := target - num

		if j, ok := numMap[diff]; ok {
			return []int{j, i}
		}

		numMap[num] = i
	}

	return nil
}
