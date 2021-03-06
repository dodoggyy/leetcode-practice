package medium

// Use LCS:
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 13 ms, faster than 36.73%
// Memory Usage: 7 MB, less than 38.78%
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range word1 {
		for j := range word2 {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// LCS concept: lengthM + lengthN - 2*LCS = change times
	return m + n - 2*dp[m][n]
}
