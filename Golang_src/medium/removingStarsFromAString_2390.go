package medium

// Use stack:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 26 ms, faster than 77.78%
// Memory Usage: 7.9 MB, less than 37.4%
func removeStars(s string) string {
	stack := []byte{}
	for i := range s {
		switch s[i] {
		case '*':
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}
