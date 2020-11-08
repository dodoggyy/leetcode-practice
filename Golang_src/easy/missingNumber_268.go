package easy

// Use Bit Manipulation:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 12 ms, faster than 99.34%
// Memory Usage: 6 MB, less than 100.00%
func missingNumber(nums []int) int {
	size := len(nums)
	result := 0

	for i := 0; i < size; i++ {
		result ^= i
		result ^= nums[i]
	}

	result ^= size

	return result
}
