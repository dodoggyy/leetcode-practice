package easy

// Use Bit Manipulation:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.1 MB, less than 100.00%
func isPowerOfFour(num int) bool {
	return num > 0 && num&(num-1) == 0 && num&0x55555555 != 0
}
