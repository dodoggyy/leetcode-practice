package medium

// Use DP:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 266 ms, faster than 52.94%
// Memory Usage: 20.4 MB, less than 76.47%
func mostPoints(questions [][]int) int64 {
	// if skip question i+1:
	//f[i+1]=max(f[i+1],f[i])

	// if not skip:
	// f[j] = max(f[j],f[i]+point[i])

	size := len(questions)
	dp := make([]int, size+1)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i, v := range questions {
		dp[i+1] = max(dp[i+1], dp[i])
		j := i + v[1] + 1
		if j > size {
			j = size
		}
		dp[j] = max(dp[j], dp[i]+v[0])
	}

	return int64(dp[size])
}
