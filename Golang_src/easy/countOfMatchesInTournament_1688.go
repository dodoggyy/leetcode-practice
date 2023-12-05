package easy

// Use Simulate:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.98 MB, less than 37.50%
func numberOfMatches(n int) int {
	result := 0
	for n > 0 {
		result += n / 2
		n = n/2 + n%2
		if n == 1 {
			break
		}
	}

	return result
}

// Use math:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 2 ms, faster than 16.41%
// Memory Usage: 1.98 MB, less than 37.50%
func numberOfMatches2(n int) int {
	return n - 1
}
