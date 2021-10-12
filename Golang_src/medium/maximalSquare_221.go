package medium

// Use dp:
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 4.6 MB, less than 18.05%
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	result := 0

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				result = 1
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			if dp[i][j] > result {
				result = dp[i][j]
			}
		}
	}

	return result * result
}
