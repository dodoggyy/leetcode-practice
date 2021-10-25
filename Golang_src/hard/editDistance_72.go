package hard

// Use DP:
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 16 ms, faster than 89.02%
// Memory Usage: 6.7 MB, less than 99.19%
func minDistance(word1 string, word2 string) int {
	size1, size2 := len(word1), len(word2)

	dp := make([][]int, size1+1)
	for i := range dp {
		dp[i] = make([]int, size2+1)
	}

	for i := 1; i <= size1; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= size2; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= size1; i++ {
		for j := 1; j <= size2; j++ {
			if word1[i-1] != word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i-1][j-1]
			}
			dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
			dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
		}
	}

	return dp[size1][size2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
