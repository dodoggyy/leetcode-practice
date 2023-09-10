package easy

// Use hashmap:
// Time Complexity: O(n))
// Space Complexity:O(n)
// Runtime: 6 ms, faster than 27.38%
// Memory Usage: 3.13 MB, less than 33.33%
func maximumNumberOfStringPairs(words []string) int {
	result := 0

	hashmap := map[string]bool{}

	revert := func(word string) string {
		str := []byte(word)
		for i := 0; i < len(str)/2; i++ {
			str[i], str[len(str)-1-i] = str[len(str)-1-i], str[i]
		}
		return string(str)
	}

	for _, word := range words {
		if hashmap[word] {
			result++
		}
		hashmap[revert(word)] = true
	}

	return result
}
