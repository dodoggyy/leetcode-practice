package medium

// Use greedy:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 93.77%
// Memory Usage: 4.2 MB, less than 100.00%
func canJump(nums []int) bool {
	reach := 0
	size := len(nums)

	for i := 0; i < size; i++ {
		if i > reach {
			return false
		}
		if i+nums[i] > reach {
			reach = i + nums[i]
		}
	}
	return true
}

// Use greedy:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 93.77%
// Memory Usage: 4.2 MB, less than 100.00%
func canJump2(nums []int) bool {
	size := len(nums)
	lastPos := size - 1

	for i := size - 2; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}
