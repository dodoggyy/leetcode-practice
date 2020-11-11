package easy

// Use Bit Manipulation:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.9 MB, less than 100.00%
func hasAlternatingBits(n int) bool {
	n ^= (n >> 1)
	return n&(n+1) == 0
}
