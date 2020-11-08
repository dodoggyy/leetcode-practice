package easy

// Use Bit Manipulation:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.5 MB, less than 100.00%
func reverseBits(num uint32) uint32 {
	result := uint32(0)
	i := 0
	for i < 32 {
		result <<= 1
		result |= (num & 1)
		num >>= 1
		i++
	}

	return result
}
