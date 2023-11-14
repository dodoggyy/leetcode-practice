package medium

// Use DP:
// Time Complexity: O(m*n)
// Space Complexity:O(n)
// Runtime: 2 ms, faster than 81.05%
// Memory Usage: 2.34 MB, less than 88.63%
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, n)
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			if j > 0 && obstacleGrid[i][j-1] == 0 {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[len(dp)-1]
}
