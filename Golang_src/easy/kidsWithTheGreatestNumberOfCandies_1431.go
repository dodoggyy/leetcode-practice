package easy

import "math"

// Use iterative:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.32 MB, less than 95.02%
func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))

	maxCandies := math.MinInt32
	for _, v := range candies {
		if v > maxCandies {
			maxCandies = v
		}
	}

	for i, v := range candies {
		if v+extraCandies >= maxCandies {
			result[i] = true
		}
	}

	return result
}
