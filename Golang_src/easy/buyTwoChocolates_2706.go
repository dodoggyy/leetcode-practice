package easy

import "math"

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 9 ms, faster than 38.10%
// Memory Usage: 3.44 MB, less than 71.43%
func buyChoco(prices []int, money int) int {
	min1, min2 := math.MaxInt32, math.MaxInt32

	for _, v := range prices {
		if v < min1 {
			min2 = min1
			min1 = v
		} else if v < min2 {
			min2 = v
		}
	}

	if min1+min2 > money {
		return money
	}

	return money - min1 - min2
}
