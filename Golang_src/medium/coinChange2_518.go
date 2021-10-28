package medium

// Use dp:
// Time Complexity: O(n*amount)
// Space Complexity:O(amount)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.3 MB, less than 77.22%
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)

	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}
