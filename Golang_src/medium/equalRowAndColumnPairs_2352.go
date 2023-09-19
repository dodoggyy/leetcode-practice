package medium

import "strconv"

// Use hashmap:
// Time Complexity: O(n^2)
// Space Complexity:O(n^2)
// Runtime: 186 ms, faster than 25.00%
// Memory Usage: 59.92 MB, less than 5.24%
func equalPairs(grid [][]int) int {
	n := len(grid)
	hashmap := map[string]int{}
	result := 0

	tmp := ""
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp += strconv.Itoa(grid[i][j]) + "-"
		}
		hashmap[tmp]++
		tmp = ""
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp += strconv.Itoa(grid[j][i]) + "-"
		}
		if v, ok := hashmap[tmp]; ok {
			result += v
		}

		tmp = ""
	}

	return result
}
