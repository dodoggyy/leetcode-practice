package easy

// Given a string s, return true if a permutation of the string could form a palindrome.

// Example 1:
// Input: s = "code"
// Output: false

// Example 2:
// Input: s = "aab"
// Output: true

// Example 3:
// Input: s = "carerac"
// Output: true

// Use cnt:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.9 MB, less than 46.15%
func canPermutePalindrome(s string) bool {
	cnt := make([]int, 128)
	for i := range s {
		cnt[s[i]]++
	}

	hasOdd := false
	for _, v := range cnt {
		if v%2 != 0 {
			if hasOdd {
				return false
			} else {
				hasOdd = true
			}
		}
	}

	return true
}
