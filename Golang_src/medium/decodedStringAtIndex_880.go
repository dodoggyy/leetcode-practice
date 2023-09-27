package medium

// Use work backwards:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 1 ms, faster than 77.78%
// Memory Usage: 1.94 MB, less than 33.33%
func decodeAtIndex(s string, k int) string {
	cnt := 0
	for i := range s {
		if isDigit2(s[i]) {
			times := int(s[i] - '0')
			cnt = cnt * times
		} else {
			cnt++
		}
	}

	idx := len(s) - 1
	for isDigit2(s[idx]) || (cnt != k && idx > 0 && k > 0) {
		if isDigit2(s[idx]) {
			times := int(s[idx] - '0')
			cnt /= times

			if k != cnt {
				k %= cnt
			}
		} else {
			cnt--
		}
		idx--
	}

	return string(s[idx])
}

func isDigit2(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}
