package medium

// Use slide window:
// Time Complexity: O(n)
// Space Complexity:O(k)
// Runtime: 4 ms, faster than 90.23%
// Memory Usage: 2.6 MB, less than 100.00%
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	result := 0
	record := [128]int{}
	start := 0

	for end := 0; end < len(s); end++ {
		// start index = max( current start index, previous current char index)
		start = max(start, record[s[end]])
		result = max(result, end-start+1) // calculate cur substring length and compare result
		record[s[end]] = end + 1          // record same char position + 1
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
