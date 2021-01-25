package easy

// Use two pointers
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 16 ms, faster than 72.77%
// Memory Usage: 6.6 MB, less than 17.82%
func validPalindrome(s string) bool {
	str := []rune(s)
	left, right := 0, len(str)-1

	for left < right {
		if str[left] != str[right] {
			return validPalindromeRange(str, left+1, right) || validPalindromeRange(str, left, right-1)
		}

		left++
		right--
	}

	return true
}

func validPalindromeRange(str []rune, left, right int) bool {
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
