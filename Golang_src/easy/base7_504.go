package easy

import (
	"strconv"
)

// Use iterative:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 100.00%
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	result := ""
	isNegative := false

	if num < 0 {
		num = -num
		isNegative = true
	}

	for num != 0 {
		result = strconv.Itoa(num%7) + result
		num /= 7
	}

	if isNegative {
		result = "-" + result
	}

	return result
}

// Use recursive:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 100.00%
func convertToBase72(num int) string {
	if num < 0 {
		return "-" + convertToBase72(-num)
	}
	if num < 7 {
		return strconv.Itoa(num)
	}

	return convertToBase72(num/7) + strconv.Itoa(num%7)
}
