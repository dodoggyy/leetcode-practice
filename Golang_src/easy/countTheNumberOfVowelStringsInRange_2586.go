package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 5 ms, faster than 80.68%
// Memory Usage: 4.9 MB, less than 99.30%
func vowelStrings(words []string, left int, right int) int {
	result := 0

	isVowel := func(ch byte) bool {
		return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
	}

	for i := left; i <= right; i++ {
		word := words[i]
		if isVowel(word[0]) && isVowel(word[len(word)-1]) {
			result++
		}
	}

	return result
}
