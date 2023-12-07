package easy

// Use greedy:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 5 ms, faster than 84.42%
// Memory Usage: 6.78 MB, less than 45.45%
func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if int(num[i]-'0')%2 == 1 {
			return string(num[:i+1])
		}
	}

	return ""
}
