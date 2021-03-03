package medium

// Use DP:
// dp[i] = dp[j] && check(s[j:i])
// Time Complexity: O(n^2)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.2 MB, less than 80.27%
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := map[string]bool{}
	for _, word := range wordDict {
		wordDictSet[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
