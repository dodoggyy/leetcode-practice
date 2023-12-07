package easy

// Use math:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.98 MB, less than 31.91%
func totalMoney(n int) int {
	totalWeek := n / 7
	remain := n % 7

	return (28+(28+7*(totalWeek-1)))*totalWeek/2 +
		((1+totalWeek)+(1+totalWeek+remain-1))*remain/2
}

// (1 + 7)*7/2 = 28
// (2 + 8)*7/2 = 35
// (3 + 9)*7/2 = 42
