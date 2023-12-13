package easy

// Use Precompute the Number of Ones in each Row and Column:
// Time Complexity: O(m*n)
// Space Complexity:O(m+n)
// Runtime: 16 ms, faster than 41.18%
// Memory Usage: 6.18 MB, less than 100.00%
func numSpecial(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	result := 0
	sumRow, sumCol := make([]int, m), make([]int, n)
	for i := range mat {
		for j, v := range mat[i] {
			sumRow[i] += v
			sumCol[j] += v
		}
	}
	for i := range mat {
		for j, v := range mat[i] {
			if v == 1 && sumRow[i] == 1 && sumCol[j] == 1 {
				result++
			}
		}
	}
	return result
}
