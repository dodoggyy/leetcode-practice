package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 3 ms, faster than 42.25%
// Memory Usage: 2.7 MB, less than 36.62%
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	str1, str2 := "", ""

	for _, v := range word1 {
		str1 += v
	}

	for _, v := range word2 {
		str2 += v
	}

	return str1 == str2
}
