package easy

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 3.8 MB, less than 14.59%
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charNum := [26]int{}

	tmpS := []rune(s)
	tmpT := []rune(t)

	for i := range tmpS {
		charNum[tmpS[i]-'a']++
		charNum[tmpT[i]-'a']--
	}

	for _, v := range charNum {
		if v != 0 {
			return false
		}
	}

	return true
}
