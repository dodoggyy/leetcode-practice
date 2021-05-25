package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 20.95%
func plusOne(digits []int) []int {
	size := len(digits)
	digits[size-1] += 1

	// O(n)
	for i := size - 1; i > 0; i-- {
		digits[i-1] += digits[i] / 10
		digits[i] %= 10
	}

	if digits[0] >= 10 {
		digits[0] %= 10
		return append([]int{1}, digits...)
	}

	return digits
}
