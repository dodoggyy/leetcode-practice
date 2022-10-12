package easy

import "sort"

// Use sort + greedy:
// Time Complexity: O(nlogn)
// Space Complexity:O(1)
// Runtime: 104 ms, faster than 16.16%
// Memory Usage: 6.7 MB, less than 71.72%
func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			return nums[i-2] + nums[i-1] + nums[i]
		}
	}

	return 0
}
