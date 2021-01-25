package medium

// Use backtracking:
// S: all result of sum
// Time Complexity: O(S), upper bound: O(n*2^n)
// Space Complexity:O(target)
// Runtime: 4 ms, faster than 72.36%
// Memory Usage: 3.1 MB, less than 57.82%
func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	var dfs func(index int, tmp []int, sum int)

	dfs = func(index int, tmp []int, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			newTmp := make([]int, len(tmp))
			copy(newTmp, tmp)
			result = append(result, newTmp)
			return
		}

		for i := index; i < len(candidates); i++ {
			tmp = append(tmp, candidates[i])
			// i for element can repeat, if change to i+1 cannot
			dfs(i, tmp, sum+candidates[i])
			// pruning
			tmp = tmp[:len(tmp)-1]
		}
	}
	// reset to previous node
	dfs(0, []int{}, 0)

	return result
}
