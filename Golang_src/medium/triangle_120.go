package medium

import "math"

// Use DP:
// Time Complexity: O(n^2)
// Space Complexity:O(n^2)
// Runtime: 4 ms, faster than 80.80%
// Memory Usage: 3.8 MB, less than 27.17%
func minimumTotal(triangle [][]int) int {
	size := len(triangle)
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
	}

	dp[0][0] = triangle[0][0]
	for i := 1; i < size; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	result := math.MaxInt32

	for _, v := range dp[size-1] {
		if v < result {
			result = v
		}
	}

	return result
}

// Use Optimize DP:
// Time Complexity: O(n^2)
// Space Complexity:O(n^2)
// Runtime: 11 ms, faster than 26.81%
// Memory Usage: 3.3 MB, less than 82.97%
func minimumTotal2(triangle [][]int) int {
	size := len(triangle)
	dp := make([]int, size)

	dp[0] = triangle[0][0]

	for i := 1; i < size; i++ {
		// need reverse traverse to prevent cover.
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp[j] = min(dp[j], dp[j-1]) + triangle[i][j]
		}
		dp[0] += triangle[i][0]
	}

	result := math.MaxInt32

	for _, v := range dp {
		if v < result {
			result = v
		}
	}

	return result
}
