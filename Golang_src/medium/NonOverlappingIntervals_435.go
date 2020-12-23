package medium

import (
	"math"
	"sort"
)

// Use greedy algorithm:
// Time Complexity: O(nlogn)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 59.02%
// Memory Usage: 4.1 MB, less than 26.23%
func eraseOverlapIntervals(intervals [][]int) int {
	if intervals == nil || len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	result := 0
	end := math.MinInt32
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			end = intervals[i][1]
		} else {
			result++
		}
	}
	return result
}
