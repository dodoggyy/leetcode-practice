package easy

// Use KMP algo:
// Time Complexity: O(m+n)
// Space Complexity:O(n)
// Runtime: 1 ms, faster than 78.52%
// Memory Usage: 2.00 MB, less than 50.20%
func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	j := 0
	next := make([]int, n)

	getNext := func() {
		j := 0
		next[0] = j
		for i := 1; i < len(needle); i++ {
			for j > 0 && needle[i] != needle[j] {
				j = next[j-1]
			}
			if needle[i] == needle[j] {
				j++
			}
			next[i] = j
		}
	}

	getNext()

	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}

	return -1
}
