package basic

/*
	https://www.hackerrank.com/challenges/one-week-preparation-kit-lonely-integer/problem

	Count Sorting - Each time a value occurs in the original array,
	you increment the counter at that index.

	For this exercise, always return a frequency array with 100 elements.

	-------------------------
	We need to initialize a frequency array of size 100 with default values of 0.

	Iterate over arr.
		[1, 1, 3, 2, 1]
		[0, 3, 1, 1, 0]
		- For every element in arr, use the element as the index in the frequency
		array and increment the count at that index by 1.

	Return the frequency array.
*/

func CountingSort(arr []int32) []int32 {
	// Initialize our frequency array
	freq := make([]int32, 100)

	// Iterate and increment our frequency array
	for _, num := range arr {
		freq[num]++
	}

	return freq
}
