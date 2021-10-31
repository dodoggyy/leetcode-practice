package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 72.73%
// Memory Usage: 3.6 MB, less than 100.00%
func finalValueAfterOperations(operations []string) int {
	result := 0

	for _, str := range operations {
		switch str {
		case "++X", "X++":
			result++
		case "--X", "X--":
			result--
		}
	}

	return result
}
