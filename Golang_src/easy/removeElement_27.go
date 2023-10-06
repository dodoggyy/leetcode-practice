package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.16 MB, less than 13.61%
func removeElement(nums []int, val int) int {
	idx := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums[i], nums[idx] = nums[idx], nums[i]
			idx--
		}
	}
	return idx + 1
}
