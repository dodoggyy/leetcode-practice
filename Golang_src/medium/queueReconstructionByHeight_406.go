package medium

import "sort"

// Use greedy algorithm:
// Time Complexity: O(nlogn + n^2) = O(n^2)
// Space Complexity:O(logn)
// Runtime: 12 ms, faster than 91.14%
// Memory Usage: 6 MB, less than 43.04%
func reconstructQueue(people [][]int) [][]int {
	// asceding order by height
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	result := make([][]int, len(people))

	for _, person := range people {
		space := person[1] + 1
		for i := range result {
			if result[i] == nil {
				space--
				if space == 0 {
					result[i] = person
					break
				}
			}
		}
	}

	return result
}
