package medium

// Use greedy algorithm
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.1 MB, less than 67.47%
func partitionLabels(S string) []int {
	last := [26]int{}
	for index, char := range S {
		last[char-'a'] = index
	}
	result := []int{}
	start, end := 0, 0
	for index := range S {
		end = max(end, last[S[index]-'a'])
		if index == end {
			result = append(result, end-start+1)
			start = end + 1
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
