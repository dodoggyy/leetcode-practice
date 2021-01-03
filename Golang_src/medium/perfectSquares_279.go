package medium

import "math"

// Use DP:
// Time Complexity: O(n^2)
// Space Complexity:O(n)
// Runtime: 32 ms, faster than 66.29%
// Memory Usage: 5.8 MB, less than 81.46%
func numSquares(n int) int {
	// numSquares(n)=min(numSquares(n-k) + 1) ?k?square numbers
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt32
		for j := 1; i-j*j >= 0; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Use Lagrange's four-square theorem:
// Time Complexity: O(n^(1/2))
// Space Complexity:O(1)
// Runtime: 32 ms, faster than 66.29%
// Memory Usage: 2 MB, less than 97.75%
func numSquares2(n int) int {
	// n = 4^k * (8*m + 7)
	for n%4 == 0 {
		n /= 4
	}
	if n%8 == 7 {
		return 4
	}

	if isSquare(n) {
		return 1
	}

	// enumeration to check if the number can be decomposed into sum of two squares.
	for i := 1; i*i <= n; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}

	// bottom case of three-square theorem.
	return 3
}

func isSquare(n int) bool {
	square := int(math.Sqrt(float64(n)))
	return n == square*square
}
