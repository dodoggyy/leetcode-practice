package easy

// Use Reduce to Smaller Problem:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 20 ms, faster than 98.96%
// Memory Usage: 6.3 MB, less than 34.38%
func checkPossibility(nums []int) bool {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			count++
			if i-2 >= 0 && nums[i-2] > nums[i] {
				nums[i] = nums[i-1]
			} else {
				nums[i-1] = nums[i]
			}
		}
	}

	return count <= 1
}
