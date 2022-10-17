package easy

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.1 MB, less than 87.25%
func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}
	hashmap := map[rune]bool{}
	cnt := 0

	for _, v := range sentence {
		if _, ok := hashmap[v]; !ok {
			cnt++
		}
		hashmap[v] = true
		if cnt == 26 {
			return true
		}
	}

	return false
}
