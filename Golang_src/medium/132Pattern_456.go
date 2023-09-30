package medium

import "math"

// Use monotonic stack:
// Time Complexity: O(n^2)
// Space Complexity:O(logn)
// Runtime: 64 ms, faster than 52.50%
// Memory Usage: 12.71 MB, less than 40.00%
func find132pattern(nums []int) bool {
	stack := []int{}
	size := len(nums)
	k := math.MinInt32

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// iterative 1 value
	for i := size - 1; i >= 0; i-- {
		if nums[i] < k {
			return true
		}
		// stack hold 3
		// k hold 2
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			k = max(k, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}

	return false
}
