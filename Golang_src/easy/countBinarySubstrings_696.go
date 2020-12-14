package easy

// Use linear scan:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 100.00%
// Memory Usage: 5.9 MB, less than 27.27%
func countBinarySubstrings(s string) int {
	count := 0
	preLen, curLen := 0, 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curLen++
		} else {
			count += min(preLen, curLen)
			preLen, curLen = curLen, 1
		}
	}

	return count + min(preLen, curLen)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
