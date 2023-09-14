package medium

import "sort"

// Use sort + two pointers:
// Time Complexity: O(nlogn + n)
// Space Complexity:O(1)
// Runtime: 109 ms, faster than 73.78%
// Memory Usage: 8.26 MB, less than 83.53%
func maxOperations(nums []int, k int) int {
	result := 0

	sort.Ints(nums)

	l, r := 0, len(nums)-1
	for l < r {
		cur := nums[l] + nums[r]
		if cur == k {
			result++
			l++
			r--
		} else if cur < k {
			l++
		} else {
			r--
		}
	}

	return result
}
