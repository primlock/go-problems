// https://leetcode.com/problems/top-k-frequent-elements/description/

package array

func TopKFrequent(nums []int, k int) []int {
	// Make a frequency map of nums
	fmap := make(map[int]int)
	for _, val := range nums {
		fmap[val]++
	}

	// Create buckets for each nums frequency
	buckets := make([][]int, len(nums)+1)
	for num, freq := range fmap {
		buckets[freq] = append(buckets[freq], num)
	}

	topK := make([]int, 0)

	// Iterate backward over buckets until len(result) == K
	for i := len(buckets) - 1; i >= 0; i-- {
		for _, num := range buckets[i] {
			topK = append(topK, num)
			if len(topK) == k {
				return topK
			}
		}
	}

	return nil
}
