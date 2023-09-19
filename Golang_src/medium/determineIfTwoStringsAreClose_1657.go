package medium

// Use hashmap:
// Time Complexity: O(m+n)
// Space Complexity:O(m+n)
// Runtime: 86 ms, faster than 20.41%
// Memory Usage: 7.26 MB, less than 28.16%
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	hashmap1, hashmap2 := map[byte]int{}, map[byte]int{}
	for i := range word1 {
		hashmap1[word1[i]]++
	}
	for i := range word2 {
		hashmap2[word2[i]]++
	}

	hashmap3, hashmap4 := map[int]int{}, map[int]int{}
	for k, v := range hashmap1 {
		if _, ok := hashmap2[k]; !ok {
			return false
		}
		hashmap3[v]++
	}
	for _, v := range hashmap2 {
		hashmap4[v]++
	}

	for k, v := range hashmap3 {
		if hashmap4[k] != v {
			return false
		}
	}

	return true
}
