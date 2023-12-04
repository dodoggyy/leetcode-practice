package easy

import (
	"strconv"
	"strings"
)

// Use string comparison:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 2 ms, faster than 61.90%
// Memory Usage: 2.20 MB, less than 66.67%
func largestGoodInteger(num string) string {
	for i := 9; i >= 0; i-- {
		str := strings.Repeat(strconv.Itoa(i), 3)
		if strings.Contains(num, str) {
			return str
		}
	}
	return ""
}
