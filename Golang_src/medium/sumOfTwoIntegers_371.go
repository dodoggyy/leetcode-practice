package medium

// Use Bit Manipulation via iterative:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.9 MB, less than 100.00%
func getSum(a int, b int) int {
	for b != 0 {
		carry := a & b << 1
		a = a ^ b
		b = carry
	}

	return a
}

// Use Bit Manipulation via recursive:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 31.94%
func getSum2(a int, b int) int {
	if b == 0 {
		return a
	}

	return getSum2((a ^ b), ((a & b) << 1))
}
