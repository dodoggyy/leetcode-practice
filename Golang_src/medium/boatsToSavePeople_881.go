package medium

import "sort"

// Use sort + two pointers:
// Time Complexity: O(nlogn + n) = O(nlogn)
// Space Complexity:O(logn)
// Runtime: 86 ms, faster than 72.50%
// Memory Usage: 8.1 MB, less than 22.50%
func numRescueBoats(people []int, limit int) int {
	result := 0
	sort.Ints(people)
	l, r := 0, len(people)-1

	for l <= r {
		cur := people[l] + people[r]
		if cur <= limit {
			result++
			l++
			r--
		} else {
			result++
			r--
		}
	}

	return result
}
