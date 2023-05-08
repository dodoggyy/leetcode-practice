package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 7.89 ms, faster than 7.89%
// Memory Usage: 5.6 MB, less than 26.32%
func diagonalSum(mat [][]int) int {
	size := len(mat)
	result := 0

	for i := 0; i*i <= (size-1)*(size-1); i++ {
		result += mat[i][i]
		result += mat[i][size-i-1]
	}

	if size%2 != 0 {
		result -= mat[size/2][size/2]
	}

	return result
}
