package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 3 ms, faster than 86.84%
// Memory Usage: 3.6 MB, less than 34.21%
func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) <= 2 {
		return true
	}

	calRate := func(i, j int) float32 {
		if coordinates[i][1]-coordinates[j][1] == 0 {
			return 1
		}
		return float32((coordinates[i][0] - coordinates[j][0])) / (float32(coordinates[i][1] - coordinates[j][1]))
	}

	rate := calRate(0, 1)
	for i := 1; i < len(coordinates)-1; i++ {
		if calRate(i, i+1) != rate {
			return false
		}
	}

	return true
}
