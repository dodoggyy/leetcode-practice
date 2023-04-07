package medium

import "math"

// Use two pointers:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 133 ms, faster than 35.80%
// Memory Usage: 13.7 MB, less than 64.20%
func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
	result := math.MaxInt32

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	idxPre := -1
	for i := range wordsDict {
		if wordsDict[i] == word1 || wordsDict[i] == word2 {
			if idxPre != -1 && (wordsDict[idxPre] != wordsDict[i] || word1 == word2) {
				result = min(result, i-idxPre)
			}
			idxPre = i
		}
	}

	return result
}
