/*
New Year's Day is around the corner and the department store is having a sale.
They have a list of items they are considering but they may need to remove some
of them. Determine the minimum number of items to remove from an array of prices
so that the sum of prices of any k items does not exceed a threshold.

Note: If the number of items in the list is less than k, then there is no need
to remove any more items.

Example:
prices = [3, 2, 1, 4, 6, 5]
k = 3
threshold = 14

The sum of prices for every k = 3 items must not be more than threshold = 14.
The sum of the prices of the last three items is 6 + 5 + 4 = 15. The item priced
at $6 can be removed leaving:

prices = [3, 2, 1, 4, 5]

No 3 items' prices sum to greater than 14. Only 1 item needs to be removed.

Function Description:
Complete the function reduceGifts in the editor below.

reduceGifts has the following parameters:
	prices int32[]: the prices of each item
	k int32: the number of items to sum
	threshold int32: the maximum price of k items

Returns:
int32: the minimum number of items to remove from the list.

Constraints:
* 1 <= k <= n <= 10^5
* 1 <= threshold <= 10^9
* 1 <= prices[i] <= 10^9

*/

package other

// prices = [3, 2, 1, 4, 6, 5]

// 5
// prices = [6, 7, 1, 12, 2, 3, 4]
// k = 3
// threshold = 12

func ReduceGifts(prices []int, k, threshold int) int {
	if k > len(prices) {
		return 0
	}

	n := len(prices)

	removed_items := 0
	windowPrices := 0

	// Initialize the window
	for i := 0; i < k; i++ {
		windowPrices += prices[i]
		if windowPrices > threshold {
			removed_items++
		}
	}

	// Slide the window to the end
	for i := k; i < n; i++ {
		windowPrices += prices[i] - prices[i-k]
		if windowPrices > threshold {
			removed_items++
		}
	}

	return removed_items
}
