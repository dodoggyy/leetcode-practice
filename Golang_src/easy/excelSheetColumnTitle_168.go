package easy

// Use iterative:
// Time Complexity: O(logn)
// Space Complexity:O(logn)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 0.0 MB, less than 100%
func convertToTitle(n int) string {
	result := ""

	for n != 0 {
		n--
		result = string('A'+(n%26)) + result
		n /= 26
	}

	return result
}

// Use recursive:
// Time Complexity: O(logn)
// Space Complexity:O(logn)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.9 MB, less than 100.00%
func convertToTitle2(n int) string {
	if n == 0 {
		return ""
	}
	n--
	return convertToTitle2(n/26) + string('A'+(n%26))
}
