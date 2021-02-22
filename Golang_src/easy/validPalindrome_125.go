package easy

import "strings"

// Use two pointer:
// Time Complexity: O(len(s))
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 64.55%
// Memory Usage: 2.9 MB, less than 69.95%
func isPalindrome(s string) bool {
	size := len(s)
	left, right := 0, size-1
	s = strings.ToLower(s)

	for left <= right {
		for left < size && isAlphabet(s[left]) != true {
			left++
		}
		for right >= 0 && isAlphabet(s[right]) != true {
			right--
		}
		if left >= size || right < 0 {
			break
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphabet(ch byte) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}

// Use optimize two pointer:
// Time Complexity: O(len(s))
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 64.55%
// Memory Usage: 2.9 MB, less than 69.95%
func isPalindrome2(s string) bool {
	size := len(s)
	left, right := 0, size-1
	s = strings.ToLower(s)

	for left < right {
		for left < right && !isAlphabet2(s[left]) {
			left++
		}
		for left < right && !isAlphabet2(s[right]) {
			right--
		}
		if left > right {
			break
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphabet2(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9'
}
