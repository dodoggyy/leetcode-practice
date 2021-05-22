package medium

import "strconv"

// Use iterative:
// Time Complexity: O(m*n + (m+n) + (m+n)) = O(m*n)
// Space Complexity:O(m+n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 3 MB, less than 43.27%
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	result := make([]int, m+n)

	// O(m*n)
	for i := m - 1; i >= 0; i-- {
		curDigit1 := int(num1[i] - '0')
		for j := n - 1; j >= 0; j-- {
			curDigit2 := int(num2[j] - '0')
			result[i+j+1] += curDigit1 * curDigit2
		}
	}

	// O(m+n)
	for i := m + n - 1; i > 0; i-- {
		result[i-1] += result[i] / 10
		result[i] %= 10
	}
	resultStr := ""
	idx := 0
	if result[0] == 0 {
		idx = 1
	}
	// O(m+n)
	for ; idx < m+n; idx++ {
		resultStr += strconv.Itoa(result[idx])
	}

	return resultStr
}
