package easy

import "sort"

// Use sorting:
// Time Complexity: O(mlogm)
// Space Complexity:O(m)
// Runtime: 11 ms, faster than 61.84%
// Memory Usage: 4.9 MB, less than 72.37%
func kWeakestRows(mat [][]int, k int) []int {
	m := len(mat)

	arr := make([]int, m)
	idx := make([]int, m)

	for i := range mat {
		cnt := 0
		for j := range mat[i] {
			cnt += mat[i][j]
		}
		arr[i] = cnt
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		a, b := arr[idx[i]], arr[idx[j]]
		return a < b || (a == b && idx[i] < idx[j])
	})

	return idx[:k]
}
