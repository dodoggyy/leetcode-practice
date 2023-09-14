package easy

// Use math:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.21 MB, less than 8.42%
func largestAltitude(gain []int) int {
	result := 0

	for i := range gain {
		if i > 0 {
			gain[i] += gain[i-1]
		}
		if result < gain[i] {
			result = gain[i]
		}
	}

	return result
}
