package medium

// Use DP:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 100.00%
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, cur, tmp := 1, 1, 0
	for i := 1; i < len(s); i++ {
		switch {
		case s[i] == '0' && s[i-1] != '1' && s[i-1] != '2':
			return 0
		// dp[i] = dp[i-2]
		case s[i] == '0':
			cur = pre
		//dp[i] = dp[i-1] + dp[i-2], dp[i-1] = dp[i-2]
		case (s[i] <= '6' && (s[i-1] == '1' || s[i-1] == '2')) || (s[i] > '6' && s[i-1] == '1'):
			tmp = cur
			cur += pre
			pre = tmp
		// dp[i] = dp[i-1]
		default:
			pre = cur
		}
	}
	return cur
}
