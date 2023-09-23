package medium

// Use dfs:
// Time Complexity: O(n+m)
// Space Complexity:O(n)
// Runtime: 6 ms, faster than 67.06%
// Memory Usage: 4.1 MB, less than 91.27%
func canVisitAllRooms(rooms [][]int) bool {
	visit := make([]bool, len(rooms))

	var dfs func(int)
	dfs = func(room int) {
		visit[room] = true
		keys := rooms[room]
		for _, key := range keys {
			if visit[key] {
				continue
			}
			dfs(key)
		}
	}

	dfs(0)
	for _, v := range visit {
		if !v {
			return false
		}
	}
	return true
}
