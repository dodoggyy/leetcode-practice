package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 2 ms, faster than 77.78%
// Memory Usage: 2.50 MB, less than 17.46%
func minOperations(s string) int {
	cnt := 0

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := range s {
		if i%2 != int(s[i]-'0') {
			cnt++
		}
	}

	return min(cnt, len(s)-cnt)
}
