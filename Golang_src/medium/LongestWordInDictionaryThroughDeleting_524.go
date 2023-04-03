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
	result := ""
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

// Use two pointers 2:
// Time Complexity: O(d*(m+n))
// Space Complexity:O(1)
// Runtime: 12 ms, faster than 93.75%
// Memory Usage: 8.1 MB, less than 25%
func findLongestWord2(s string, dictionary []string) string {
	result := ""
	for _, t := range dictionary {
		i := 0
		for j := range s {
			if s[j] == t[i] {
				i++
				if i == len(t) {
					if len(t) > len(result) || len(t) == len(result) && t < result {
						result = t
					}
					break
				}
			}
		}
	}

	return result
}
