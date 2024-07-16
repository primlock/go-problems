// https://leetcode.com/problems/valid-anagram/description/

package leetcode

import "reflect"

func ValidAnagram(s, t string) bool {
	sMap := make(map[rune]int)
	for _, char := range s {
		sMap[char] += 1
	}

	tMap := make(map[rune]int)
	for _, char := range t {
		tMap[char] += 1
	}

	return reflect.DeepEqual(sMap, tMap)
}
