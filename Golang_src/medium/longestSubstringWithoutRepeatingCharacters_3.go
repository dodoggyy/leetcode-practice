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

// Use slide window2:
// Time Complexity: O(n)
// Space Complexity:O(k)
// k for character set
// Runtime: 3 ms, faster than 90.23%
// Memory Usage: 2.7 MB, less than 100.00%
func lengthOfLongestSubstring2(s string) int {
	arr := make([]int, 128)
	result := 0

	l, r := 0, 0

	for r < len(s) {
		idx := s[r]
		l = max(l, arr[idx])
		result = max(result, r-l+1)

		arr[idx] = r + 1
		r++
	}

	return result
}
