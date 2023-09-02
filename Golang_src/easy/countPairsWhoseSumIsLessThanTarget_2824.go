package easy

import "sort"

// Use two pointer:
// Time Complexity: O(nlogn + n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.67 MB, less than 60.10%
func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	result := 0

	l, r := 0, len(nums)-1
	for l < r {
		cur := nums[l] + nums[r]
		if cur < target {
			result += r - l
			l++
		} else {
			r--
		}
	}

	return result
}
