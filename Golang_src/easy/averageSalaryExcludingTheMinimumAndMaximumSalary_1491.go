package easy

import "math"

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.7 MB, less than 53.60%
func average(salary []int) float64 {
	min, max := math.MaxInt64, math.MinInt64
	cnt := 0

	for _, v := range salary {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		cnt += v
	}

	return float64(cnt-min-max) / float64(len(salary)-2)
}
