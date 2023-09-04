package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 6 ms, faster than 37.71%
// Memory Usage: 2.68 MB, less than 95.86%
func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	result := 0
	for _, v := range hours {
		if v >= target {
			result++
		}
	}

	return result
}
