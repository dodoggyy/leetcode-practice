package medium

import "sort"

// Use greedy algorithm:
// Time Complexity: O(nlogn)
// Space Complexity:O(1)
// Runtime: 76 ms, faster than 36.00%
// Memory Usage: 6.9 MB, less than 58.00%
func findMinArrowShots(points [][]int) int {
	if points == nil || len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	result := 1
	end := points[0][1]

	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			result++
			end = points[i][1]
		}
	}

	return result
}
