package medium

import "sort"

// Use backtracking:
// Time Complexity: O(n*n!)
// Space Complexity:O(n)
// Runtime: 4 ms, faster than 90.57%
// Memory Usage: 3.8 MB, less than 91.98%
func permuteUnique(nums []int) [][]int {
	result := [][]int{}

	sort.Ints(nums) // O(nlogn)

	size := len(nums)
	keep := []int{}
	visit := make([]bool, size)

	var dfs func()
	dfs = func() {
		if len(keep) == size {
			tmp := make([]int, size)
			copy(tmp, keep)
			result = append(result, tmp)
			return
		}

		for i := range nums {
			if visit[i] || i > 0 && nums[i] == nums[i-1] && !visit[i-1] {
				continue
			}
			visit[i] = true
			keep = append(keep, nums[i])
			dfs()
			visit[i] = false
			keep = keep[:len(keep)-1]
		}
	}

	dfs()

	return result
}
