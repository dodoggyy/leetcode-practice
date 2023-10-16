package easy

// Use recursive:
// Time Complexity: O(rowIndex^2)
// Space Complexity:O(rowIndex)
// Runtime: 1 ms, faster than 77.04%
// Memory Usage: 2.10 MB, less than 35.35%
func getRow(rowIndex int) []int {
	result := make([][]int, rowIndex+1)

	for i := range result {
		result[i] = make([]int, i+1)
		result[i][0], result[i][i] = 1, 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}

	return result[rowIndex]
}

// Use optimize recursive:
// Time Complexity: O(rowIndex^2)
// Space Complexity:O(1)
// Runtime: 1 ms, faster than 77.04%
// Memory Usage: 2.02 MB, less than 35.35%
func getRow2(rowIndex int) []int {
	var pre, cur []int
	for i := 0; i <= rowIndex; i++ {
		cur = make([]int, i+1)
		cur[0], cur[i] = 1, 1
		for j := 1; j < i; j++ {
			cur[j] = pre[j-1] + pre[j]
		}
		pre = cur
	}
	return cur
}

// Use math:
// Time Complexity: O(rowIndex)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.94 MB, less than 79.15%
func getRow3(rowIndex int) []int {
	// C(m,n) = n! / (m!*(n-m)!)
	result := make([]int, rowIndex+1)
	result[0] = 1
	for i := 1; i <= rowIndex; i++ {
		result[i] = result[i-1] * (rowIndex - i + 1) / i
	}

	return result
}
