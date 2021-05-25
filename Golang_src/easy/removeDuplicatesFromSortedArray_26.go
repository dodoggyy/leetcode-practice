package easy

// Use two pointers:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 88.17%
// Memory Usage: 4.6 MB, less than 100.00%
func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	result := 1
	idx := 0
	cur := nums[0]

	// O(n)
	for i := 1; i < len(nums); i++ {
		if nums[i] != cur {
			cur = nums[i]
			result++
			idx++
			nums[idx], nums[i] = nums[i], nums[idx]
		}
	}

	return result
}
