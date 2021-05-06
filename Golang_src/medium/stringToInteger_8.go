package medium

import "math"

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.3 MB, less than 100.00%
func myAtoi(s string) int {
	result := 0
	idx := 0
	isNeg := false
	size := len(s)

	// whitespace ignore
	for _, ch := range s {
		if ch != ' ' {
			break
		}
		idx++
	}

	// null data handle
	if idx == size {
		return result
	}

	// negtive handle
	if s[idx] == '-' {
		isNeg = true
	}
	if s[idx] == '-' || s[idx] == '+' {
		idx++
	}

	// traversal digit
	for idx < size && isDigits(s[idx]) {
		curNum := int(s[idx] - '0')

		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && curNum > 7) {
			if isNeg {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
		result = result*10 + curNum
		idx++
	}

	if isNeg {
		return -result
	}

	return result
}

func isDigits(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}
