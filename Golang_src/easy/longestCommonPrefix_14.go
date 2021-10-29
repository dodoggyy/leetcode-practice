package easy

// Use scan:
// Time Complexity: O(n*s)
// Space Complexity:O(1)
// n: strings average length
// s: string counts
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.4 MB, less than 100.00%
func longestCommonPrefix(strs []string) string {
	result := strs[0]

	idx := 0

	for idx < len(result) {
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) <= idx || result[idx] != strs[i][idx] {
				return result[:idx]
			}
		}
		idx++
	}

	return result
}
