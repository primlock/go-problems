/*
https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem

- We want to know how many times Maria breaks her season record for MOST and LEAST points
in a game.
- The first game establishes the record for the season (the baseline).
- Return an array of size 2 where index 0 is for breaking MOST point records and index 1 is
for breaking LEAST point records.
---------------------------------------------------
We need to keep a tally of the number of time we update min and max.
- We also need to handle the base case of the first game setting the baseline.

For each iteration:
- Is this value a new max?
	- Yes = increment the counter, update the max
	- No  = next iteration
- Is this value a new min?
	- Yes = increment the counter, update the min
	- No  = next iteration

Handling first game case:
	- use the index var in range for i == 0.

After all iterations, create and return a new slice with MIN counter and MAX counter.
- return []int{min_count, max_count}
*/

package basic

func BreakingRecords(scores []int32) []int32 {
	var min_counter int32
	var max_counter int32
	var min_points int32
	var max_points int32

	for i, points := range scores {
		if i == 0 {
			min_points = points
			max_points = points
			continue
		}

		if points < min_points {
			min_counter++
			min_points = points
		} else if points > max_points {
			max_counter++
			max_points = points
		}
	}

	return []int32{max_counter, min_counter}
}
