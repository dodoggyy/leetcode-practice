package medium

// Use DP:
// Time Complexity: O(n^2)
// Space Complexity:O(n)
// Runtime: 64 ms, faster than 86.00%
// Memory Usage: 2 MB, less than 40.00%
func integerBreak(n int) int {
	dp := make([]int, n+1)

	dp[2] = 1
	for i := 3; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = max(curMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}

	return dp[n]
}

// Use optimize DP:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 64 ms, faster than 86.00%
// Memory Usage: 2 MB, less than 40.00%
func integerBreak2(n int) int {
	dp := make([]int, n+1)

	dp[2] = 1
	for i := 3; i <= n; i++ {
		// only nedd consider divide to 2 and 3 case, >= 4 cannot greater than 2 or 3
		dp[i] = max(max(2*(i-2), 2*dp[i-2]), max(3*(i-3), 2*dp[i-3]))
	}

	return dp[n]
}
