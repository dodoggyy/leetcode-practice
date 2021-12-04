package easy

// Use iterative:
// Time Complexity: O(numRows^2)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.1 MB, less than 99.43%
func generate(numRows int) [][]int {
	result := [][]int{}

	for i := 0; i < numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0], tmp[i] = 1, 1

		for j := 1; j < i; j++ {
			tmp[j] = result[i-1][j-1] + result[i-1][j]
		}

		result = append(result, tmp)
	}

	return result
}
