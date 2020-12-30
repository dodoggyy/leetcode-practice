package easy

// Use greedy algorithm:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 95.92%
// Memory Usage: 3.3 MB, less than 100.00%
func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	preSum, result := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if preSum < 0 {
			preSum = nums[i]
		} else {
			preSum += nums[i]
		}
		result = max(result, preSum)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Use DP:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 95.92%
// Memory Usage: 3.3 MB, less than 100.00%
func maxSubArray2(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		result = max(result, nums[i])
	}

	return result
}
