package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 90.99%
// Memory Usage: 4.79 MB, less than 99.14%
func sortArrayByParity(nums []int) []int {
	cur := 0

	for i := range nums {
		if nums[i]%2 == 0 {
			nums[i], nums[cur] = nums[cur], nums[i]
			cur++
		}
	}

	return nums
}
