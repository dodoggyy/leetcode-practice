package medium

// Use brute force:
// Time Complexity: O(n^3)
// Space Complexity:O(1)
// Runtime: 164 ms, faster than 31.64%
// Memory Usage: 2.7 MB, less than 48.91%
func longestPalindrome(s string) string {
	size := len(s)
	result := ""
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			if j-i+1 > len(result) && isPalindrome(s[i:j+1]) {
				result = s[i : j+1]
			}
		}
	}
	return result
}

func isPalindrome(s string) bool {
	size := len(s)
	for i := 0; i < size/2+1; i++ {
		if s[i] != s[size-i-1] {
			return false
		}
	}
	return true
}

// Use DP:
// Time Complexity: O(n^2)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 0.0 MB, less than 100%
func longestPalindrome2(s string) string {
	size := len(s)
	length := 0
	start := 0

	getLen := func(s string, l, r int) int {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		// due to out of range of l & r,
		// currently we need (l+1) ~ (r-1)
		return (r - 1) - (l + 1) + 1
	}

	for i := 0; i < size; i++ {
		cur := max(getLen(s, i, i), getLen(s, i, i+1))
		if cur > length {
			length = cur
			start = i - (cur-1)/2
		}
	}

	return s[start : start+length]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
