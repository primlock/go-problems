/*
There is a new mathmatical game for kids. You will be given n sticks and the player
is required to form rectangles from those sticks.

Formally, given an array of n integers representing the lengths of the sticks, you
are required to create rectangles using those sticks. Note that a particular stick
can be used in at most one rectangle and in order to create a rectangle we must have
exactly two pairs of sticks with the same lengths. For example, you can create a rectangle
using sticks of lengths [2, 2, 5, 5] and [4, 4, 4, 4,] but not with [3, 3, 5, 8]. The
goal of the game is to maximize the total sum of the areas of all the rectangles
formed.

In order to make the game more interesting, we are allowed to reduce any integer by
at most 1. Given the array sideLengths, representing the length of the sticks, find the
maximum sum of areas of rectangles that can be formed such that each element of the array
can be used as length or breadth of at more one rectangle and you are allowed to decrease
any integer by at most 1. Since this number can be quite large, return the answer modulo
10^9+7.

Note: It is not a requirement that all side lengths be used. Also, a modulo b here represents
the remainder obtained when an integer a is divided by an integer b.

Example:
The side lengths are given as sideLengths = [2, 6, 6, 2, 3, 5].

The lengths 2, 2, 6 and 6 can be used to form a rectangle of area 2 * 6 = 12. No other rectangles
can be formed with the remaining lengths. The answer is 12 modulo (10^9 + 7) = 12.

Function Description:
Complete the function getMaxTotalArea in the editor below.
getMaxTotalArea has the following parameters(s):
	sideLengths int32[]: the side lengths that can be used to form rectangles.

Returns:
int32: the maximum total area of the rectangles that can be formed, modulo (10^9 + 7).

Constraints:
* 1 <= n <= 10^5
* 2 <= sideLengths[i] <= 10^6 where 0 <= i < n

*/

package other

func GetMaxTotalArea(sideLengths []int32) int32 {
	return 0
}
