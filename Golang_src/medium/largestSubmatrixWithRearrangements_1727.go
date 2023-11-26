package medium

import "sort"

// Use sorting:
// Time Complexity: O(m*n)
// Space Complexity:O(1)
// Runtime: 124 ms, faster than 100.00%
// Memory Usage: 15.13 MB, less than 100.00%
func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	result := 0
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				matrix[i][j] += matrix[i-1][j]
			}
		}
	}
	for i := 0; i < m; i++ {
		sort.Ints(matrix[i])
		for j := n - 1; j >= 0; j-- {
			height := matrix[i][j]
			if result < height*(n-j) {
				result = height * (n - j)
			}
		}
	}
	return result
}
