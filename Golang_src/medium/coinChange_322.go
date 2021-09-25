package medium

import "math"

// Use DP:
// dp[i] = min(dp[i- coinVal] + 1, dp[i])
// Time Complexity: O(n*s)
// Space Complexity:O(n)
// n : amount
// s: coin type
// Runtime: 8 ms, faster than 94.75%
// Memory Usage: 6.5 MB, less than 96.35%
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, v := range coins {
			if v <= i {
				dp[i] = min(dp[i-v]+1, dp[i])
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
