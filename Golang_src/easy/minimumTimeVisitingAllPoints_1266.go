package easy

// Use math:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 3.00 MB, less than 59.09%
func minTimeToVisitAllPoints(points [][]int) int {
	result := 0

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 1; i < len(points); i++ {
		diffX := abs(points[i][0] - points[i-1][0])
		diffY := abs(points[i][1] - points[i-1][1])

		result += min(diffX, diffY) + abs(diffX-diffY)
	}

	return result
}
