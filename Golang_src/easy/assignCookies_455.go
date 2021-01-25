package easy

import "sort"

// Use greedy algorithm:
// Time Complexity: O(g*log(g) + s*log(s))
// Space Complexity:O(1)
// Runtime: 40 ms, faster than 18.46%
// Memory Usage: 6.1 MB, less than 13.85%
func findContentChildren(g []int, s []int) int {
	if g == nil || s == nil {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)

	childIndex, cookieIndex := 0, 0

	for childIndex < len(g) && cookieIndex < len(s) {
		if g[childIndex] <= s[cookieIndex] {
			childIndex++
		}
		cookieIndex++
	}

	return childIndex
}
