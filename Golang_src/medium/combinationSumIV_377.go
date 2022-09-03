package medium

// Use DP:
// Time Complexity: O(target*n)
// Space Complexity:O(target)
// Runtime: 3 ms, faster than 33.93%
// Memory Usage: 2 MB, less than 81.23%
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, v := range nums {
			if i-v >= 0 {
				dp[i] += dp[i-v]
			}
		}
	}

	return dp[target]
}
