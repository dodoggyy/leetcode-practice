package medium

// Use linear search:
// Time Complexity: O(m+n)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 16.59%
// Memory Usage: 2.7 MB, less than 77.02%
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	curM, curN := 0, 0

	for curM < m && matrix[curM][curN] <= target {
		curM++
	}

	if curM == 0 {
		curM++
	}

	for curN < n && matrix[curM-1][curN] <= target {
		curN++
	}

	if curN == 0 {
		curN++
	}

	return matrix[curM-1][curN-1] == target
}

// Use binary search:
// Time Complexity: O(log(mn))
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.6 MB, less than 100.00%
func searchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	left, right := 0, m*n-1
	for left <= right {
		idx := left + (right-left)/2
		curVal := matrix[idx/n][idx%n]

		if curVal == target {
			return true
		} else if curVal < target {
			left = idx + 1
		} else {
			right = idx - 1
		}
	}

	return false
}
