package easy

// Use copy directly:
// Time Complexity: O(r*c)
// Space Complexity:O(r*c)
// Runtime: 9 ms, faster than 52.24%
// Memory Usage: 5.96 MB, less than 100.00%
func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, m)
	}

	for i := range matrix {
		for j := range matrix[i] {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}
