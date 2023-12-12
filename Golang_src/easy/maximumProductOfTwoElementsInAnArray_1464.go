package easy

import "math"

// Use copy directly:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 2 ms, faster than 84.00%
// Memory Usage: 2.60 MB, less than 57.00%
func maxProduct(nums []int) int {
	max1, max2 := math.MinInt32, math.MinInt32
	for _, v := range nums {
		if v > max1 {
			max2 = max1
			max1 = v
		} else if v > max2 {
			max2 = v
		}
	}

	return (max1 - 1) * (max2 - 1)
}
