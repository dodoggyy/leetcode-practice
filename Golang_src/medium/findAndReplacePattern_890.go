package medium

// Use hashmap:
// n: words list length
// m: pattern length
// Time Complexity: O(n*m)
// Space Complexity:O(m)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.3 MB, less than 64.29%
func findAndReplacePattern(words []string, pattern string) []string {
	match := func(word, pattern string) bool {
		hashmap := map[byte]byte{}
		for i := range word {
			if val, ok := hashmap[word[i]]; ok {
				if val != pattern[i] {
					return false
				}
			} else {
				hashmap[word[i]] = pattern[i]
			}
		}

		return true
	}

	result := []string{}

	for _, word := range words {
		if match(word, pattern) && match(pattern, word) {
			result = append(result, word)
		}
	}

	return result
}
