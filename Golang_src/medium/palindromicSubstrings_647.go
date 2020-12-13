package medium

// Use Expand Around Center:
// Time Complexity: O(n^2)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 52.52%
func countSubstrings(s string) int {
	count := 0
	strS := []rune(s)
	size := len(strS)

	for i := 0; i < size; i++ {
		helper(strS, i, i, &count)
		helper(strS, i, i+1, &count)
	}

	return count
}

func helper(strS []rune, indexStart, indexEnd int, count *int) {
	for indexStart >= 0 && indexEnd < len(strS) && strS[indexStart] == strS[indexEnd] {
		indexStart--
		indexEnd++
		*count++
	}
}
