package medium

import (
	"strconv"
	"strings"
)

// Use backtracking:
// Time Complexity: O(3^SEGMENT_COUNT * |s|)
// Space Complexity:O(SEGMENT_COUNT)
// SEGMENT_COUNT: 4
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.3 MB, less than 52.70%
func restoreIpAddresses(s string) []string {
	result := []string{}
	size := len(s)
	if size < 4 || size > 12 {
		return result
	}

	keep := []string{}

	var dfs func(start int)
	dfs = func(start int) {
		if len(keep) == 4 && start == size {
			tmp := make([]string, len(keep))
			copy(tmp, keep)
			result = append(result, strings.Join(tmp, "."))
			return
		}
		// check temp result >= 4 && current index not equal size
		if len(keep) == 4 && start < size {
			return
		}

		for i := 1; i <= 3; i++ {
			if start+i > size {
				return
			}
			// check start with 0 cannot be leading digit
			if i != 1 && s[start] == '0' {
				return
			}
			str := s[start : start+i]
			// check digit <= 255 or not
			if n, _ := strconv.Atoi(str); n > 255 {
				return
			}
			keep = append(keep, str)
			dfs(start + i)
			keep = keep[:len(keep)-1]
		}
	}

	dfs(0)

	return result
}
