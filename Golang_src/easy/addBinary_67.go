package easy

import "strconv"

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.7 MB, less than 53.60%
func addBinary(a string, b string) string {
	result := ""

	strA := []rune(a)
	strB := []rune(b)

	indexA := len(strA) - 1
	indexB := len(strB) - 1

	carry := 0

	for indexA >= 0 || indexB >= 0 || carry == 1 {
		if indexA >= 0 {
			if strA[indexA] == '1' {
				carry++
			}
			indexA--
		}
		if indexB >= 0 {
			if strB[indexB] == '1' {
				carry++
			}
			indexB--
		}
		result = strconv.Itoa(carry%2) + result
		carry /= 2
	}

	return result
}
