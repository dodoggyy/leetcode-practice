package easy

// Use iterative:
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 8 ms, faster than 95.38%
// Memory Usage: 6.5 MB, less than 36.97%
func matrixReshape(mat [][]int, r int, c int) [][]int {
	result := make([][]int, r)
	for i := range result {
		result[i] = make([]int, c)
	}

	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}

	for i := 0; i < m*n; i++ {
		result[i/c][i%c] = mat[i/n][i%n]
	}

	return result
}
