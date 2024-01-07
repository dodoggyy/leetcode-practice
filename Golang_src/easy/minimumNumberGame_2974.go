package easy

import "sort"

// Use sorting:
// Time Complexity: O(nlogn)
// Space Complexity:O(1)
// Runtime: 7 ms, faster than 47.72%
// Memory Usage: 2.89 MB, less than 85.28%
func numberGame(nums []int) []int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i += 2 {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}

	return nums
}
