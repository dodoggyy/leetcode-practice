package medium

// Use 0-1 backpack:
// Time Complexity: O(n*target)
// Space Complexity:O(target)
// Runtime: 4 ms, faster than 93.18%
// Memory Usage: 2 MB, less than 100%
func numRollsToTarget(n int, k int, target int) int {
	if target < n || n*k < target {
		return 0
	}

	// prevent overflow
	// 10^9 + 7 is prime number
	mod := int(1e9 + 7)

	dp := make([]int, target+1)
	dp[0] = 1

	// traverse item
	for i := 0; i < n; i++ {
		// traverse backpack
		for j := target; j >= 0; j-- {
			dp[j] = 0

			// traverse decision
			for m := 1; m <= k; m++ {
				if j >= m {
					dp[j] = (dp[j] + dp[j-m]) % mod
				}
			}
		}
	}

	return dp[target]
}
