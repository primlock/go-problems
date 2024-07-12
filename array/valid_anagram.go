// https://leetcode.com/problems/valid-anagram/description/

package leetcode

import "reflect"

func ValidAnagram(s, t string) bool {
	sMap := make(map[string]int)
	for _, char := range s {
		sMap[string(char)] += 1
	}

	tMap := make(map[string]int)
	for _, char := range t {
		tMap[string(char)] += 1
	}

	return reflect.DeepEqual(sMap, tMap)
}
