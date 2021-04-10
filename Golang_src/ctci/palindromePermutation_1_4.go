package ctci

import "unicode"

// Use hash table:
// Time Complexity: O(n)
// Space Complexity:O(n)
func palindromePermutation(str string) bool {
	hashmap := map[rune]int{}

	for _, ch := range str {
		unicode.ToLower(ch)
		if ch < 'a' || ch > 'z' {
			continue
		}
		if val, exist := hashmap[ch]; exist {
			hashmap[ch] = 1
		} else {
			hashmap[ch] = val + 1
		}
	}

	hasOdd := false
	for _, val := range hashmap {
		if (val % 2) == 1 {
			if hasOdd {
				return false
			} else {
				hasOdd = true
			}
		}
	}

	return true
}
