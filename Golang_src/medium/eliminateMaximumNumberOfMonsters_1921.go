package medium

import "sort"

// Use sort:
// Time Complexity: O(nlogn)
// Space Complexity:O(n)
// Runtime: 137 ms, faster than 46.67%
// Memory Usage: 9.51 MB, less than 53.33%
func eliminateMaximum(dist []int, speed []int) int {
	size := len(dist)

	timeArrival := make([]float64, size)
	for i := range dist {
		timeArrival[i] = float64(dist[i])/float64(speed[i])
	}
	sort.Float64s(timeArrival)
	for i := range timeArrival {
		if timeArrival[i] <= float64(i) {
			return i
		}
	}
	return size
}
