package medium

import "sort"

// Use sorting + sliding window:
// Time Complexity: O(nolgn)
// Space Complexity:O(logn)
// Runtime: 200 ms, faster than 34.21%
// Memory Usage: 9.78 MB, less than 18.42%
func maxFrequency(nums []int, k int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	result := 1
	sort.Ints(nums)
	total := 0

	for l, r := 0, 1; r < len(nums); r++ {
		total += (nums[r] - nums[r-1]) * (r - l)
		for total > k {
			total -= nums[r] - nums[l]
			l++
		}
		result = max(result, r-l+1)
	}

	return result
}
