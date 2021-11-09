package medium

// Use DP:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 76.92%
// Memory Usage: 2.7 MB, less than 100.00%
func maxProduct(nums []int) int {
	maxP, minP, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		maxTmp, minTmp := maxP, minP
		maxP = max(maxTmp*nums[i], max(nums[i], minTmp*nums[i]))
		minP = min(minTmp*nums[i], min(nums[i], maxTmp*nums[i]))

		result = max(result, maxP)
	}

	return result
}
