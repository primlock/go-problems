package basic

import "fmt"

/*
	https://www.hackerrank.com/challenges/apple-and-orange/problem

	Print the number of APPLES and number of ORANGES that will land on the house.
	The trees and the house are on a line, the x-axis.

	House - located in the range [s, t]:
	- s = start point
	- t = end point

	Trees - each located at a single point:
	- a = apple tree
	- b = orange tree

	Fruit - lands d units away from the tree on the x-axis:
	- positive d = landed to the right
	- negative d = landed to the left
	- m = apples
	- n = oranges

	Find how many apples and oranges land in the range [s, t]
	-----------------------------------------------------------
	The location of the fruit can be found by taking the sum of (d + a) for apples and (d + b) for oranges.
	- For apples: if sum is gte to s and lte to t, it landed on the house
	- For oranges: if sum is gte to s and lte to t, it landed on the house.

	Keep an individual counter for apples and oranges.
*/

func CountApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	apple_counter := 0
	orange_counter := 0

	for _, apple := range apples {
		// Did the apple travel far enough?
		if (apple+a) >= s && (apple+a) <= t {
			apple_counter++
		}
	}

	for _, orange := range oranges {
		// Did the orange travel far enough?
		if (orange+b) >= s && (orange+b) <= t {
			orange_counter++
		}
	}

	fmt.Println(apple_counter)
	fmt.Println(orange_counter)
}
