package medium

// Use Layer-by-Layer:
// Time Complexity: O(m*n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 100.00%
func spiralOrder(matrix [][]int) []int {
	row, column := len(matrix), len(matrix[0])
	left, right := 0, column-1
	top, bottom := 0, row-1

	result := []int{}

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		for j := top + 1; j <= bottom; j++ {
			result = append(result, matrix[j][right])
		}
		if left < right && top < bottom {
			for k := right - 1; k > left; k-- {
				result = append(result, matrix[bottom][k])
			}
			for m := bottom; m > top; m-- {
				result = append(result, matrix[m][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}

	return result
}

// Use Layer-by-Layer 2:
// Time Complexity: O(m*n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.1 MB, less than 52.82%
func spiralOrder2(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, n-1
	top, bottom := 0, m-1

	result := []int{}

	for left <= right && top <= bottom {
		for i := left; i < right; i++ {
			result = append(result, matrix[top][i])
		}
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}

		// consider last center may not have cases as below
		if left < right && top < bottom {
			for i := right - 1; i > left; i-- {
				result = append(result, matrix[bottom][i])
			}

			for i := bottom; i > top; i-- {
				result = append(result, matrix[i][left])
			}
		}

		left++
		right--
		top++
		bottom--
	}

	return result
}
