package medium

import (
	"strconv"
	"strings"
)

// Use stack:
// Time Complexity: O(S)
// Space Complexity:O(S)
// S: string length
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 81.37%
func decodeString(s string) string {
	stackNum := []int{}
	stackStr := []string{}
	num := 0
	result := ""

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			n, _ := strconv.Atoi(string(ch))
			num = num*10 + n
		} else if ch == '[' {
			stackStr = append(stackStr, result)
			result = ""
			stackNum = append(stackNum, num)
			num = 0
		} else if ch == ']' {
			str := stackStr[len(stackStr)-1]
			stackStr = stackStr[:len(stackStr)-1]
			count := stackNum[len(stackNum)-1]
			stackNum = stackNum[:len(stackNum)-1]

			result = str + strings.Repeat(result, count)
		} else {
			result += string(ch)
		}
	}
	return result
}
