package medium

import "sort"

// Use sorting:
// Time Complexity: O(nlogn)
// Space Complexity:O(logn)
// Runtime: 106 ms, faster than 88.00%
// Memory Usage: 8.11 MB, less than 72.00%
func maxCoins(piles []int) int {
	sort.Ints(piles)
	result := 0
	l, r := 0, len(piles)
	for l < r {
		r -= 2
		l++
		result += piles[r]
	}

	return result
}

// 1 2 2 4 7 8

// 1 2 3 4 5 6 7 8 9

// 1 [8] 9
// 2 [6] 7
// 3 [4] 5
