package ctci

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 0.0 MB, less than 100%
func IsOneAway(s1, s2 string) bool {
	size1, size2 := len(s1), len(s2)

	if abs(size1, size2) > 1 {
		return false
	}

	if size1 == size2 { // replace
		return oneEditReplace(s1, s2)
	} else if size1 > size2 { // insert
		return oneEditInsert(s2, s1)
	} else { // delete
		// size1 < size2
		return oneEditInsert(s1, s2)
	}
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func oneEditReplace(s1, s2 string) bool {
	isDiff := false
	for i := range s1 {
		if s1[i] != s2[i] {
			if isDiff {
				return false
			} else {
				isDiff = true
			}
		}
	}
	return true
}

func oneEditInsert(s1, s2 string) bool {
	idx1, idx2 := 0, 0

	for idx1 < len(s1) && idx2 < len(s2) {
		if s1[idx1] != s2[idx2] {
			if idx1 != idx2 {
				return false
			} else {
				idx2++
			}
		} else {
			idx1++
			idx2++
		}
	}
	return true
}
