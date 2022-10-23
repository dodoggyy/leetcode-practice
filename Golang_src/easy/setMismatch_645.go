package easy

// Use Constant Space:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 48 ms, faster than 57.41%
// Memory Usage: 6.8 MB, less than 44.44%
func findErrorNums(nums []int) []int {
	for i := range nums {
		for nums[i] != (i+1) && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != (i + 1) {
			return []int{nums[i], i + 1}
		}
	}
	return []int{}
}
