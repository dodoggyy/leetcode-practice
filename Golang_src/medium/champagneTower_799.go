package medium

// Use simulation:
// Time Complexity: O(query_row^2)
// Space Complexity:O(query_row)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 4.88 MB, less than 48.8%
func champagneTower(poured int, query_row int, query_glass int) float64 {
	row := []float64{float64(poured)}

	for i := 1; i <= query_row; i++ {
		rowNext := make([]float64, i+1)
		for j := range row {
			if row[j] > 1 {
				rowNext[j] += (row[j] - 1) / 2
				rowNext[j+1] += (row[j] - 1) / 2
			}
		}
		row = rowNext
	}

	if row[query_glass] > 1 {
		return 1
	}

	return row[query_glass]
}
