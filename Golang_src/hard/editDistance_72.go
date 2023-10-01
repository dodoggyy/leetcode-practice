package hard

// Use DP:
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 5.55 MB, less than 79.87%
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := range dp {
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}

	for i, ch1 := range word1 {
		for j, ch2 := range word2 {
			if ch1 == ch2 {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j], min(dp[i+1][j], dp[i][j+1])) + 1
			}

		}
	}

	return dp[m][n]
}
