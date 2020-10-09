package medium

import "strings"

// Use two pointers:
// Time Complexity: O(n*x)
// Space Complexity:O(x)
// n for number of string list d
// x for max string in list d
// Runtime: 16 ms, faster than 68.97%
// Memory Usage: 6.4 MB, less than 89.66%
func findLongestWord(s string, d []string) string {
	var result string = ""
	target := []rune(s)

	for _, str := range d {
		tmp := []rune(str)

		if checkSubStr(target, tmp) {
			if len(tmp) > len(result) || (len(tmp) == len(result) && strings.Compare(str, result) < 0) {
				result = string(tmp)
			}
		}
	}

	return result
}

func checkSubStr(target, str []rune) bool {
	i, j := 0, 0

	for i < len(target) && j < len(str) {
		if str[j] == target[i] {
			j++
		}
		i++
	}

	return j == len(str)
}
