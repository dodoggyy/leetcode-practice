package easy

import "strings"

// Use brute force:
// Time Complexity: O(k*n^2)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.69 MB, less than 100.00%
func countPrefixSuffixPairs(words []string) int {
	result := 0

	isPair := func(x, y string) bool {
		return strings.HasPrefix(x, y) && strings.HasSuffix(x, y)
	}

	for i := 1; i < len(words); i++ {
		for j := 0; j < i; j++ {
			if isPair(words[i], words[j]) {
				result++
			}
		}
	}

	return result
}
