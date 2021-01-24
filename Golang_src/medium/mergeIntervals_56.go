package medium

import "sort"

// Use greedy:
// Time Complexity: O(nlogn)
// Space Complexity:O(logn)
// Runtime: 8 ms, faster than 92.17%
// Memory Usage: 4.7 MB, less than 91.77%
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}

	pre := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if pre[1] >= cur[0] { // overlap
			pre[1] = max(pre[1], cur[1])
		} else {
			result = append(result, pre)
			pre = cur
		}
	}

	result = append(result, pre)

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
