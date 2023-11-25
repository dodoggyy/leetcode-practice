package medium

// Use prefix sum:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 100 ms, faster than 86.21%
// Memory Usage: 8.88 MB, less than 37.93%
func getSumAbsoluteDifferences(nums []int) []int {
	size := len(nums)
	result := make([]int, size)

	sumPrefix := make([]int, size)
	sumPrefix[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sumPrefix[i] = sumPrefix[i-1] + nums[i]
	}

	for i := range nums {
		l := (i+1)*nums[i] - sumPrefix[i]
		r := sumPrefix[size-1] - sumPrefix[i] - nums[i]*(size-1-i)
		result[i] = l + r
	}

	return result
}
