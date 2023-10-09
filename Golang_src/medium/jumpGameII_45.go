package medium

// Use greedy:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 9 ms, faster than 93.88%
// Memory Usage: 6.00 MB, less than 89.80%
func jump(nums []int) int {
	end := 0
	result := 0
	maxPosition := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, nums[i]+i)
		if i == end {
			result++
			end = maxPosition
		}
	}

	return result
}
