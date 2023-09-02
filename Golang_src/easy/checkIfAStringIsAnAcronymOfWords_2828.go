package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 5 ms, faster than 53.18%
// Memory Usage: 3.6 MB, less than 100.00%
func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}

	for i := range words {
		if words[i][0] != s[i] {
			return false
		}
	}
	return true
}
