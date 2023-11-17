package medium

import "sort"

// Use sorting + greedy:
// Time Complexity: O(nolgn)
// Space Complexity:O(logn)
// Runtime: 243 ms, faster than 54.05%
// Memory Usage: 8.61 MB, less than 54.05%
func minPairSum(nums []int) int {
	result := 0
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	for l < r {
		cur := nums[l] + nums[r]
		if cur > result {
			result = cur
		}
		l++
		r--
	}

	return result
}
