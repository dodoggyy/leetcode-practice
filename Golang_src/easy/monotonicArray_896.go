package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 140 ms, faster than 55.32%
// Memory Usage: 8.6 MB, less than 89.36%
func isMonotonic(nums []int) bool {
	isIncrease, isDecrease := true, true
	for i := 1; i < len(nums); i++ {
			if nums[i-1] < nums[i] {
					isDecrease = false
			}
			if nums[i-1] > nums[i] {
					isIncrease = false
			}
	}

	return isDecrease || isIncrease
}