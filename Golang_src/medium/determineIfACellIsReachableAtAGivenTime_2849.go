package medium

// Use math:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 2 ms, faster than 35.07%
// Memory Usage: 2.62 MB, less than 100.00%
func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
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

	if sx == fx && sy == fy {
		return t != 1
	}

	distX, distY := abs(sx-fx), abs(sy-fy)

	return abs(distX-distY) <= t-min(distX, distY)
}
