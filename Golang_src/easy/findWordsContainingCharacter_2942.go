package easy

// Use iterative:
// Time Complexity: O(m*n)
// Space Complexity:O(1)
// Runtime: 7 ms, faster than 100.00%
// Memory Usage: 4.64 MB, less than 100.00%
func findWordsContaining(words []string, x byte) []int {
	result := []int{}

	for i, word := range words {
		for j := range word {
			if word[j] == x {
				result = append(result, i)
				break
			}
		}
	}

	return result
}
