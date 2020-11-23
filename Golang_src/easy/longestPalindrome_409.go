package easy

// Use Greedy:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.4 MB, less than 8.51%
func longestPalindrome(s string) int {
	charNum := [128]int{}
	tmpS := []rune(s)
	result := 0

	for _, v := range tmpS {
		charNum[v]++
	}

	for _, v := range charNum {
		result += ((v >> 1) << 1)
	}

	if result < len(tmpS) {
		result++
	}

	return result
}
