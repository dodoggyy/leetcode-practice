package medium

// Use dfs:
// Time Complexity: O(n^2)
// Space Complexity:O(n)
// Runtime: 19 ms, faster than 97.14%
// Memory Usage: 6.59 MB, less than 100.00%
func findCircleNum(isConnected [][]int) int {
	result := 0

	visited := make([]bool, len(isConnected))

	var dfs func(idx int)
	dfs = func(idx int) {
		visited[idx] = true
		for i, connect := range isConnected[idx] {
			if connect == 1 && !visited[i] {
				dfs(i)
			}
		}
	}

	for i, v := range visited {
		if !v {
			result++
			dfs(i)
		}
	}

	return result
}

// 1 1 0
// 1 1 0
// 0 0 1
