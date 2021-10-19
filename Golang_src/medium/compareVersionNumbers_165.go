package medium

// Use two pointers:
// Time Complexity: O(m + n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 44.59%
func compareVersion(version1 string, version2 string) int {
	size1, size2 := len(version1), len(version2)

	cur1, cur2 := 0, 0

	isDigit := func(ch byte) bool {
		if ch >= '0' && ch <= '9' {
			return true
		}
		return false
	}

	for cur1 < size1 || cur2 < size2 {
		tmp1, tmp2 := 0, 0
		for cur1 < size1 && isDigit(version1[cur1]) {
			tmp1 = tmp1*10 + int(version1[cur1]-'0')
			cur1++
		}
		for cur2 < size2 && isDigit(version2[cur2]) {
			tmp2 = tmp2*10 + int(version2[cur2]-'0')
			cur2++
		}
		cur1++
		cur2++

		if tmp1 < tmp2 {
			return -1
		} else if tmp1 > tmp2 {
			return 1
		}
	}

	return 0
}
