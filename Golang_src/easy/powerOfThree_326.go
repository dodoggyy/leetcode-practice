package easy

// Use recursive:
// Time Complexity: O(log3n)
// Space Complexity:O(1)
// Runtime: 16 ms, faster than 93.65%
// Memory Usage: 6.1 MB, less than 100.00%
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	for n > 1 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return true
}

// Use Integer Limitations:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 36 ms, faster than 35.71%
// Memory Usage: 6.1 MB, less than 100.00%
func isPowerOfThree2(n int) bool {
	return n > 0 && 1162261467%n == 0
}
