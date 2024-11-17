// https://leetcode.com/problems/fizz-buzz/

package array

import "fmt"

func FizzBuzz(n int) []string {
	answer := make([]string, 0)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			answer = append(answer, "FizzBuzz")
		} else if i%3 == 0 {
			answer = append(answer, "Fizz")
		} else if i%5 == 0 {
			answer = append(answer, "Buzz")
		} else {
			answer = append(answer, fmt.Sprintf("%d", i))
		}
	}

	return answer
}
