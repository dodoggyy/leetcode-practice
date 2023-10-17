package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 1 ms, faster than 78.77%
// Memory Usage: 2.12 MB, less than 32.16%
func lengthOfLastWord(s string) int {
	l, r := 0, len(s)-1

	for r >= 0 && s[r] == ' ' {
		r--
	}
	l = r
	for l >= 0 && s[l] != ' ' {
		l--
	}

	return r - l
}
