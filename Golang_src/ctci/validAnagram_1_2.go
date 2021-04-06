package ctci

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// same as LC 242. Valid Anagram
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashmap := map[byte]int{}

	for i := 0; i < len(s); i++ {
		hashmap[s[i]]++
		hashmap[t[i]]--
	}

	for _, val := range hashmap {
		if val != 0 {
			return false
		}
	}

	return true
}
