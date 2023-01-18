package medium

// Use Kadane algorithm:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 51 ms, faster than 93.55%
// Memory Usage: 7.5 MB, less than 67.74%
// Ref: https://leetcode.cn/problems/maximum-sum-circular-subarray/solution/wo-hua-yi-bian-jiu-kan-dong-de-ti-jie-ni-892u/
func maxSubarraySumCircular(nums []int) int {
	// case 1: max subarray is not circular
	// case 2: max subarray is circular

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	curMax, curMin := 0, 0
	sum := 0
	maxSum, minSum := nums[0], nums[0]

	for _, v := range nums {
		curMax = max(curMax, 0) + v
		maxSum = max(maxSum, curMax)
		curMin = min(curMin, 0) + v
		minSum = min(minSum, curMin)
		sum += v
	}
	// case 1
	if sum == minSum {
		return maxSum
	}

	// case 2
	return max(maxSum, sum-minSum)
}
