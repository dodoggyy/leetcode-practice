package easy

// Use DP:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 5.34%
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	// dp[i] = dp[i-1]+dp[i-2], for i >= 2
	// dp[0] = 1, dp[1] = 2
	dp := make([]int, n)
	dp[0], dp[1] = 1, 2

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}

// Use Fib number:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.9 MB, less than 100.00%
func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	first, second := 1, 2

	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}

	return second
}
