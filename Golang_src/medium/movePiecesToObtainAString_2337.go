package medium

// Use two pointers:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 24 ms, faster than 100.00%
// Memory Usage: 6.9 MB, less than 100.00%
func canChange(start string, target string) bool {
	if len(start) != len(target) {
		return false
	}

	size := len(start)
	idxI, idxJ := 0, 0

	for idxI < size || idxJ < size {
		for idxI < size && start[idxI] == '_' {
			idxI++
		}
		for idxJ < size && target[idxJ] == '_' {
			idxJ++
		}

		if idxI == size || idxJ == size {
			return idxI == size && idxJ == size
		}

		if start[idxI] != target[idxJ] {
			return false
		}

		if start[idxI] == 'L' && idxI < idxJ {
			return false
		}

		if start[idxI] == 'R' && idxI > idxJ {
			return false
		}

		idxI++
		idxJ++
	}

	return true
}
