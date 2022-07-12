package medium

import "sort"

// Use backtracking:
// Time Complexity: O(4^n)
// Space Complexity:O(n)
// Runtime: 215 ms, faster than 7.14%
// Memory Usage: 2.1 MB, less than 46.43%
func makesquare(matchsticks []int) bool {
	total := 0
	for _, v := range matchsticks {
		total += v
	}

	if total%4 != 0 {
		return false
	}

	sort.Ints(matchsticks)

	edges := [4]int{}

	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == -1 {
			return true
		}
		for i := range edges {
			edges[i] += matchsticks[idx]
			if edges[i] <= total/4 && dfs(idx-1) {
				return true
			}
			edges[i] -= matchsticks[idx]
		}
		return false
	}

	return dfs(len(matchsticks) - 1)
}
