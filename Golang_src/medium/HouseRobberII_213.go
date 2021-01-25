package medium

// Use DP with memorize:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.1 MB, less than 45.16%
func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	length := len(nums)
	if length == 1 {
		return nums[0]
	}

	return max(subRob(nums, 0, length-2), subRob(nums, 1, length-1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func subRob(nums []int, indexStart, indexEnd int) int {
	length := indexEnd - indexStart + 1

	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[indexStart]
	}
	dp := make([]int, length)
	dp[0] = nums[indexStart]
	dp[1] = max(nums[indexStart], nums[indexStart+1])
	for i := 2; i < length; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[indexStart+i])
	}

	return dp[length-1]
}

// Use DP without memorize:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 45.16%
func rob2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	length := len(nums)
	if length == 1 {
		return nums[0]
	}

	return max(subRob2(nums, 0, length-2), subRob2(nums, 1, length-1))
}

func subRob2(nums []int, indexStart, indexEnd int) int {
	rob, notRob := 0, 0

	for i := indexStart; i <= indexEnd; i++ {
		tmp := rob
		rob = max(rob, notRob+nums[i])
		notRob = tmp
	}
	return rob
}
