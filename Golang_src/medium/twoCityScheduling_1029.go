package medium

import "sort"

// Use sorting:
// Time Complexity: O(nlogn)
// Space Complexity:O(1)
// Runtime: 5 ms, faster than 45.83%
// Memory Usage: 2.5 MB, less than 33.33%
func twoCitySchedCost(costs [][]int) int {
	result := 0

	// sort by costA - costB
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})

	for i := range costs {
		if i < len(costs)/2 {
			result += costs[i][0]
		} else {
			result += costs[i][1]
		}
	}

	return result
}
