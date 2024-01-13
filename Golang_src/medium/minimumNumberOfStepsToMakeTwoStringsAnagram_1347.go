package medium

// Use count:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 14 ms, faster than 81.40%
// Memory Usage: 6.58 MB, less than 74.42%
func minSteps(s string, t string) int {
	result := 0

	cnt := make([]int, 26)
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for _, ch := range t {
		cur := ch - 'a'
		if cnt[cur] == 0 {
			result++
		} else {
			cnt[cur]--
		}
	}

	return result
}
