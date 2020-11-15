package medium

// Use Bit Manipulation:
// Time Complexity: O(n^2)
// Space Complexity:O(n)
// Runtime: 12 ms, faster than 91.89%
// Memory Usage: 6.1 MB, less than 13.51%
func maxProduct(words []string) int {
	result := 0
	size := len(words)
	tmp := make([]int, size)

	for i, word := range words {
		str := []rune(word)

		for _, ch := range str {
			tmp[i] |= 1 << (ch - 'a')
		}
	}

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if tmp[i]&tmp[j] == 0 {
				result = max(result, len(words[i])*len(words[j]))
			}
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
