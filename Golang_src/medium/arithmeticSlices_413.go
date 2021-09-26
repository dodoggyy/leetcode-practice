package medium

// Use DP:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 48 ms, faster than 46.22%
// Memory Usage: 2.3 MB, less than 60.44%
func numberOfArithmeticSlices(nums []int) int {
	size := len(nums)
	if size < 3 {
		return 0
	}
	dp := make([]int, size)
	result := 0

	for i := 2; i < size; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
		}
	}

	for _, v := range dp {
		result += v
	}

	return result
}

// Use diff + count:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 48 ms, faster than 46.22%
// Memory Usage: 2.3 MB, less than 60.44%
func numberOfArithmeticSlices2(nums []int) int {
	size := len(nums)
	if size < 3 {
		return 0
	}
	result := 0

	diff, tmpCnt := nums[1]-nums[0], 0
	for i := 2; i < size; i++ {
		if nums[i]-nums[i-1] == diff {
			tmpCnt++
		} else {
			diff = nums[i] - nums[i-1]
			tmpCnt = 0
		}
		result += tmpCnt
	}

	return result
}
