package medium

// Use spiral traversal:
// Time Complexity: O(n^2)
// Space Complexity:O(1)
// Runtime: 1 ms, faster than 40.52%
// Memory Usage: 2 MB, less than 79.09%
func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	left, right, up, down := 0, n-1, 0, n-1

	idx := 1
	for left <= right && up <= down {
		for i := left; i <= right; i++ {
			result[up][i] = idx
			idx++
		}
		for i := up + 1; i <= down; i++ {
			result[i][right] = idx
			idx++
		}
		if left < right && up < down {
			for i := right - 1; i >= left; i-- {
				result[down][i] = idx
				idx++
			}

			for i := down - 1; i > up; i-- {
				result[i][left] = idx
				idx++
			}
		}

		left++
		right--
		up++
		down--
	}

	return result
}
