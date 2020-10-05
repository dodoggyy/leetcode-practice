package easy

import "math"

// Use two pointers
// Time Complexity: O(logc * n^(1/2))
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.9 MB, less than 85.00%
func judgeSquareSum(c int) bool {
	if c < 0 {
		return false
	}
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		target := left*left + right*right
		if target > c {
			right--
		} else if target < c {
			left++
		} else {
			return true
		}
	}
	return false
}
