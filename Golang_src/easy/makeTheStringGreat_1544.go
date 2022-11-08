package easy

// Use stack:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.3 MB, less than 21.28%
func makeGood(s string) string {
	stack := []byte{}

	isEqual := func(a, b byte) bool {
		if a == b {
			return false
		}

		if a > b {
			return a == b+32
		}

		return a == b-32
	}

	for i := range s {
		if len(stack) > 0 && isEqual(stack[len(stack)-1], s[i]) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	result := ""
	for _, ch := range stack {
		result += string(ch)
	}

	return result
}
