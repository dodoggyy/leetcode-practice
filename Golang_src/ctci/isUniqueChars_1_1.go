package ctci

// Use iterative:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 24 ms, faster than 31.52%
// Memory Usage: 6.1 MB, less than 7.25%
func isUniqueChars(str string) bool {
	if len(str) > 128 {
		return false
	}

	cnt := make([]bool, 128)

	for _, ch := range str {
		if cnt[int(ch)] {
			return false
		}
		cnt[int(ch)] = true
	}

	return true
}
