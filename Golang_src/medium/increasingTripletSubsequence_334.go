package medium

import "math"

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 64 ms, faster than 86.00%
// Memory Usage: 17.6 MB, less than 50.00%
func increasingTriplet(nums []int) bool {
	min1, min2 := math.MaxInt32, math.MaxInt32

	for _, v := range nums {
		if v <= min1 {
			min1 = v
		} else if v <= min2 {
			min2 = v
		} else {
			return true
		}
	}
	return false
}
