package medium

// Use two pointers:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 236 ms ms, faster than 31.11%
// Memory Usage: 9.4 MB, less than 71.11%
func minCost(colors string, neededTime []int) int {
	l, r := 0, 0
	result := 0
	tmpMax := 0
	tmpResult := 0

	for r < len(colors) {
		if colors[l] == colors[r] {
			tmpResult += neededTime[r]
			tmpMax = max(tmpMax, neededTime[r])
			r++
		} else {
			if r-l > 1 {
				result += (tmpResult - tmpMax)
			}
			l = r
			tmpMax = 0
			tmpResult = 0
		}
	}

	if r-l > 1 {
		result += (tmpResult - tmpMax)
	}

	return result
}
