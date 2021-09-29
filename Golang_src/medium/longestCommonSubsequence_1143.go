package medium

// Use DP:
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 4 ms, faster than 78.57%
// Memory Usage: 11 MB, less than 88.78%
func longestCommonSubsequence(text1 string, text2 string) int {
	size1, size2 := len(text1), len(text2)

	dp := make([][]int, size1+1)

	for i := range dp {
		dp[i] = make([]int, size2+1)
	}

	for i := 1; i <= size1; i++ {
		for j := 1; j <= size2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[size1][size2]
}
