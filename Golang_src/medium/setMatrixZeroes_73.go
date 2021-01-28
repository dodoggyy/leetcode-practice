package medium

// Use Additional Memory Approach:
// Time Complexity: O(m*n)
// Space Complexity:O(m+n)
// Runtime: 12 ms, faster than 84.13%
// Memory Usage: 6 MB, less than 78.31%
func setZeroes(matrix [][]int) {
	row, col := len(matrix), len(matrix[0])

	hashmapRow := map[int]bool{}
	hashmapCol := map[int]bool{}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				hashmapRow[i] = true
				hashmapCol[j] = true
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if hashmapRow[i] || hashmapCol[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// Use the first row / first col to indicate whether the i-th row / j-th column need be zeroed:
// Time Complexity: O(m*n)
// Space Complexity:O(1)
// Runtime: 12 ms, faster than 84.13%
// Memory Usage: 5.9 MB, less than 78.31%
func setZeroes2(matrix [][]int) {
	row, col := len(matrix), len(matrix[0])
	isFirstCol := false

	for i := 0; i < row; i++ {

		// Since first cell for both first row and first column is the same i.e. matrix[0][0]
		// We can use an additional variable for either the first row/column.
		// For this solution we are using an additional variable for the first column
		if matrix[i][0] == 0 {
			isFirstCol = true
		}

		// Iterate over the array once again and using the first row and first column, update the elements.
		for j := 1; j < col; j++ {
			// If an element is zero, we set the first element of the corresponding row and column to 0
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// Iterate over the array once again and using the first row and first column, update the elements.
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// See if the first row needs to be set to zero as well
	if matrix[0][0] == 0 {
		for j := 0; j < col; j++ {
			matrix[0][j] = 0
		}
	}

	// See if the first column needs to be set to zero as well
	if isFirstCol {
		for i := 0; i < row; i++ {
			matrix[i][0] = 0
		}
	}
}
