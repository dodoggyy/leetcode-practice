package easy

// Use Reduce to Smaller Problem:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 35 ms, faster than 48.84%
// Memory Usage: 6.6 MB, less than 74.42%
func checkPossibility(nums []int) bool {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			count++
			if count > 1 {
				return false
			}
			if i > 1 && nums[i-2] > nums[i] {
				nums[i] = nums[i-1]
			} else {
				nums[i-1] = nums[i]
			}
		}
	}

	return true
}
