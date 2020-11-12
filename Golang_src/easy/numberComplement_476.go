package easy

// Use Bit Manipulation:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 100.00%
func findComplement(num int) int {
	mask := 1 << 30

	for num&mask == 0 {
		mask >>= 1
	}

	return num ^ (mask<<1 - 1)
}
