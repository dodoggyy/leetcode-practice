package easy

// Use greedy algorithm:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.1 MB, less than 21.64%
func isSubsequence(s string, t string) bool {
	tmpS, tmpT := []rune(s), []rune(t)

	index, size := 0, len(tmpS)

	if size == 0 {
		return true
	}

	for _, v := range tmpT {
		if v == tmpS[index] {
			index++
			if index == size {
				return true
			}
		}
	}

	return false
}

// Use two pointers:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 1 ms, faster than 66.70%
// Memory Usage: 1.9 MB, less than 100%
func isSubsequence2(s string, t string) bool {
	idx1, idx2 := 0, 0
	size1, size2 := len(s), len(t)

	for idx1 < size1 && idx2 < size2 {
		if s[idx1] == t[idx2] {
			idx1++
			idx2++
		} else {
			idx2++
		}
	}

	return idx1 == size1
}
