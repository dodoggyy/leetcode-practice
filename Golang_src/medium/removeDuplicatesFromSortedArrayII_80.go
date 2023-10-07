package medium

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 5 ms, faster than 56.18%
// Memory Usage: 3.04 MB, less than 32.98%
func removeDuplicates(nums []int) int {
	helper := func(times int) int {
		idx := 0
		for _, v := range nums {
			if idx < times || v != nums[idx-times] {
				nums[idx] = v
				idx++
			}
		}

		return idx
	}
	return helper(2)
}
