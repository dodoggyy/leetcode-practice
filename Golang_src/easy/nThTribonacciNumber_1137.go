package easy

// Use dp:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 1 ms, faster than 77.8%
// Memory Usage: 1.9 MB, less than 16.63%
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-3] + dp[i-2] + dp[i-1]
	}

	return dp[n]
}

// Use optimize dp:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 1 ms, faster than 77.08%
// Memory Usage: 1.89 MB, less than 97.3%
func tribonacci2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	nMinusOne, nMinusTwo, nMinusThree := 1, 1, 0

	result := 0
	for i := 3; i <= n; i++ {
		result = nMinusOne + nMinusTwo + nMinusThree
		nMinusThree = nMinusTwo
		nMinusTwo = nMinusOne
		nMinusOne = result
	}

	return result
}
