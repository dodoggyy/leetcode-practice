package easy

// Use bit manipulation:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 100.00%
func hammingDistance(x int, y int) int {
	tmp := x ^ y
	result := 0

	for tmp != 0 {
		result++
		tmp &= tmp - 1
	}

	return result
}
