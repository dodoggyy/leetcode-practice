package medium

// Use Permutation formula:
// uniquePaths(m,n) = C((m+n-2), (n-1))
// Time Complexity: O(min(m,n))
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 79.05%
func uniquePaths(m int, n int) int {
	var result int64 = 1

	s := (m - 1) + (n - 1)
	d := n - 1

	for i := 1; i <= d; i++ {
		result = result * int64(s-d+i) / int64(i)
	}

	return int(result)
}

// Use DP:
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 34.78%
func uniquePaths2(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
