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
