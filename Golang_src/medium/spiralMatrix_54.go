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
