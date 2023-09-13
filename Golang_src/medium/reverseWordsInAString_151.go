package medium

import "strings"

// Use lib:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 2 ms, faster than 71.91%
// Memory Usage: 2.88 MB, less than 93.54%
func reverseWords2(s string) string {
	strs := strings.Fields(s)
	for i := 0; i < len(strs)/2; i++ {
		strs[i], strs[len(strs)-1-i] = strs[len(strs)-1-i], strs[i]
	}

	return strings.Join(strs, " ")
}

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 6 ms, faster than 20.78%
// Memory Usage: 7.68 MB, less than 12.45%
func reverseWords(s string) string {
	tmp := []string{}
	result := ""

	reverse := func(str []byte) []byte {
		for i := 0; i < len(str)/2; i++ {
			str[i], str[len(str)-1-i] = str[len(str)-1-i], str[i]
		}

		return str
	}

	cur := len(s) - 1
	tmpStr := []byte{}
	for cur >= 0 {
		if s[cur] == ' ' {
			if len(tmpStr) > 0 {
				reverse(tmpStr)
				tmp = append(tmp, string(tmpStr))
				tmpStr = []byte{}
			}
		} else {
			tmpStr = append(tmpStr, s[cur])
		}
		cur--
	}

	if len(tmpStr) > 0 {
		reverse(tmpStr)
		tmp = append(tmp, string(tmpStr))
	}

	for _, str := range tmp {
		result += str + " "
	}

	return result[:len(result)-1]
}
