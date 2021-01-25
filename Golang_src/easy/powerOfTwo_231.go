package easy

// Use Bit Manipulation:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 42.86%
// Memory Usage: 2.2 MB, less than 100.00%
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
