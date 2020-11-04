package easy

// Use bit manipulation:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 16 ms, faster than 56.29%
// Memory Usage: 6 MB, less than 73.88%
func singleNumber(nums []int) int {
	result := 0
	for _, v := range nums {
		result ^= v
	}

	return result
}
