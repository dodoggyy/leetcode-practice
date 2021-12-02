package easy

// Use binary search:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 28 ms, faster than 91.33%
// Memory Usage: 6.9 MB, less than 68.27%
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	l, r := 0, len(nums)-1
	idx := len(result) - 1

	for l <= r {
		sqL, sqR := nums[l]*nums[l], nums[r]*nums[r]
		if sqL > sqR {
			result[idx] = sqL
			l++
		} else {
			result[idx] = sqR
			r--
		}
		idx--
	}

	return result
}
