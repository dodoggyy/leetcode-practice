package medium

import "math"

// Use sliding window:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 212 ms, faster than 66.67%
// Memory Usage: 10.1 MB, less than 33.33%
func minOperations(nums []int, x int) int {
	l, r := 0, 0
	sum := 0
	result := math.MaxInt32

	for _, v := range nums {
		sum += v
	}
	target := sum - x
	cur := 0

	for r < len(nums) {
		cur += nums[r]
		for cur > target && l <= r {
			cur -= nums[l]
			l++
		}
		if cur == target {
			result = min(result, len(nums)-(r-l+1))
		}
		r++
	}

	if result == math.MaxInt32 {
		return -1
	}

	return result
}
