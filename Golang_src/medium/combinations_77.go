package medium

// Use backtracking:
// Time Complexity: O(C(n,k)*k
// Space Complexity:O(n)
// Runtime: 12 ms, faster than 70.78%
// Memory Usage: 6.2 MB, less than 98.17%
func combine(n int, k int) [][]int {
	result := [][]int{}
	list := []int{}
	for i := 1; i <= n; i++ {
		list = append(list, i)
	}

	keep := []int{}
	visit := make([]bool, len(list))

	var dfs func(cur int)
	dfs = func(cur int) {
		if len(keep) == k {
			tmp := make([]int, len(keep))
			copy(tmp, keep)
			result = append(result, tmp)
		}
		for i := cur; i < len(list); i++ {
			if visit[i] {
				continue
			}
			visit[i] = true
			keep = append(keep, list[i])
			dfs(i + 1)
			visit[i] = false
			keep = keep[:len(keep)-1]
		}
	}

	dfs(0)

	return result
}
