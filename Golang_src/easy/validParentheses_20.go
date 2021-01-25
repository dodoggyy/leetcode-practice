package easy

// Use stack:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 47.13%
func isValid(s string) bool {
	stack := []rune{}
	isEmpty := func() bool {
		return len(stack) == 0
	}

	for _, v := range s {
		switch v {
		case '(', '{', '[':
			stack = append(stack, v)
		case ')':
			if isEmpty() {
				return false
			}
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if tmp != '(' {
				return false
			}
		case ']':
			if isEmpty() {
				return false
			}
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if tmp != '[' {
				return false
			}
		case '}':
			if isEmpty() {
				return false
			}
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if tmp != '{' {
				return false
			}
		}
	}

	return isEmpty()
}
