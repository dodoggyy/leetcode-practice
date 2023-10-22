package medium

// Use bit shift:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 7 ms, faster than 70.09%
// Memory Usage: 4.82 MB, less than 5.61%
func rangeBitwiseAnd(left int, right int) int {
	shift := 0
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	return left << shift
}
