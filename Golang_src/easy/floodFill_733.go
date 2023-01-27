package easy

// Use DFS:
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 6 ms, faster than 78.25%
// Memory Usage: 4 MB, less than 84.2%
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	m, n := len(image), len(image[0])
	org := image[sr][sc]
	if org == color {
		return image
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		image[i][j] = color
		if i-1 >= 0 && image[i-1][j] == org {
			dfs(i-1, j)
		}
		if j-1 >= 0 && image[i][j-1] == org {
			dfs(i, j-1)
		}
		if i+1 < m && image[i+1][j] == org {
			dfs(i+1, j)
		}
		if j+1 < n && image[i][j+1] == org {
			dfs(i, j+1)
		}
	}

	dfs(sr, sc)

	return image
}
