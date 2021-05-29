package medium

// Use backtracking:
// Time Complexity: O(n*n!)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.7 MB, less than 35.87%
func permute(nums []int) [][]int {
	size := len(nums)
	result := [][]int{}
	if size == 0 {
		return result
	}

	visited := map[int]bool{}

	var dfs func([]int)
	dfs = func(path []int) {
		if len(path) == size {
			tmp := make([]int, size)
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		// total: O(n*n!)
		// O(n)
		for _, num := range nums {
			if visited[num] {
				continue
			}
			path = append(path, num)
			visited[num] = true

			// O(n!)
			dfs(path)

			path = path[:len(path)-1]
			visited[num] = false
		}
	}

	dfs([]int{})

	return result
}
