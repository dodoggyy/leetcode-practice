package easy

import (
	"math"
	"sort"
)

// Use Sorting:
// Time Complexity: O(nlogn)
// Space Complexity:O(logn)
// Runtime: 56 ms, faster than 32.81%
// Memory Usage: 6.3 MB, less than 100.00%
func maximumProduct(nums []int) int {
	sort.Ints(nums)

	size := len(nums)
	result1 := nums[size-1] * nums[size-2] * nums[size-3]
	result2 := nums[0] * nums[1] * nums[size-1]

	if result1 > result2 {
		return result1
	}

	return result2
}

// Use Single Scan:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 32 ms, faster than 95.31%
// Memory Usage: 6.3 MB, less than 100.00%
func maximumProduct2(nums []int) int {
	min1, min2 := math.MaxInt64, math.MaxInt64
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64

	for _, v := range nums {

		if v <= min1 {
			min2 = min1
			min1 = v
		} else if v <= min2 {
			min2 = v
		}
		if v >= max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v >= max2 {
			max3 = max2
			max2 = v
		} else if v >= max3 {
			max3 = v
		}
	}

	result1 := min1 * min2 * max1
	result2 := max1 * max2 * max3

	if result1 > result2 {
		return result1
	}

	return result2
}
