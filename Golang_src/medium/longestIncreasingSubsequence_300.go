package medium

// Use DP:
// Time Complexity: O(n^2)
// Space Complexity:O(n)
// Runtime: 60 ms, faster than 37.96%
// Memory Usage: 3.8 MB, less than 38.81%
func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	dp := make([]int, size)
	dp[0] = 1

	result := 1

	for i := 1; i < size; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}

	return result
}
