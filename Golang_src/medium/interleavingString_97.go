package medium

// Use dp:
// Time Complexity: O(n*m)
// Space Complexity:O(n*m)
// Runtime: 2 ms, faster than 56.95%
// Memory Usage: 2.18 MB, less than 61.99%
func isInterleave(s1 string, s2 string, s3 string) bool {
	size1, size2, size3 := len(s1), len(s2), len(s3)
	if size1+size2 != size3 {
		return false
	}

	dp := make([][]bool, size1+1)
	for i := range dp {
		dp[i] = make([]bool, size2+1)
	}
	dp[0][0] = true
	for i := 0; i <= size1; i++ {
		for j := 0; j <= size2; j++ {
			cur := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[cur])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[cur])
			}
		}
	}

	return dp[size1][size2]
}

// Use optimize dp:
// Time Complexity: O(n*m)
// Space Complexity:O(m)
// Runtime: 1 ms, faster than 76.95%
// Memory Usage: 2.00 MB, less than 80.26%
func isInterleave2(s1 string, s2 string, s3 string) bool {
	size1, size2, size3 := len(s1), len(s2), len(s3)
	if size1+size2 != size3 {
		return false
	}

	dp := make([]bool, size2+1)

	dp[0] = true
	for i := 0; i <= size1; i++ {
		for j := 0; j <= size2; j++ {
			cur := i + j - 1
			if i > 0 {
				dp[j] = dp[j] && s1[i-1] == s3[cur]
			}
			if j > 0 {
				dp[j] = dp[j] || dp[j-1] && s2[j-1] == s3[cur]
			}
		}
	}

	return dp[size2]
}
