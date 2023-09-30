package medium

// Use backtracking:
// n: length of digits
// Time Complexity: O(P(9, k)) = O((9!*k)/(9-k)!)
// Space Complexity:O(K)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.91 MB, less than 70.19%
func combinationSum3(k int, n int) [][]int {
	upper := (1 + 9) * 10 / 2
	result := [][]int{}
	if k > upper {
		return result
	}

	keep := []int{}

	var dfs func(idx, remain int)
	dfs = func(idx, remain int) {
		if len(keep) == k && remain == 0 {
			tmp := make([]int, len(keep))
			copy(tmp, keep)
			result = append(result, tmp)
			return
		}

		if remain < 0 {
			return
		}

		for i := idx; i <= 9; i++ {
			keep = append(keep, i)
			dfs(i+1, remain-i)
			keep = keep[:len(keep)-1]
		}
	}

	dfs(1, n)

	return result
}
