package easy

// Use iterative:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 100.00%
func trailingZeroes(n int) int {
	if n < 0 {
		return 0
	}
	result := 0
	for n != 0 {
		n /= 5
		result += n

	}

	return result
}

// Use recursive:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 100.00%
func trailingZeroes2(n int) int {
	if n == 0 {
		return 0
	}
	return n/5 + trailingZeroes2(n/5)
}
