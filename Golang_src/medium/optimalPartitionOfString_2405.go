package medium

// Use greedy:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100%
// Memory Usage: 5.2 MB, less than 80%
func partitionString(s string) int {
	lastSeen := make([]int, 26)
	for i := range lastSeen {
		lastSeen[i] = -1
	}
	cnt, start := 1, 0

	for i := range s {
		if lastSeen[s[i]-'a'] >= start {
			cnt++
			start = i
		}
		lastSeen[s[i]-'a'] = i
	}

	return cnt
}
