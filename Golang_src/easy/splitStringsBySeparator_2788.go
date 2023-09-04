package easy

// Use iterative:
// m : length of words list
// n : length of max(word)
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 19 ms, faster than 91.82%
// Memory Usage: 7.25 MB, less than 83.64%
func splitWordsBySeparator(words []string, separator byte) []string {
	result := []string{}

	tmp := []byte{}
	for _, word := range words {
		for i := range word {
			if word[i] == separator {
				if len(tmp) > 0 {
					result = append(result, string(tmp))
					tmp = []byte{}
				}
			} else {
				tmp = append(tmp, word[i])
			}
		}
		if len(tmp) > 0 {
			result = append(result, string(tmp))
			tmp = []byte{}
		}
	}

	return result
}
