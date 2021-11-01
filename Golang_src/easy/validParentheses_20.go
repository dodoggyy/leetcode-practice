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

// Use stack 2:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 83.50%
func isValid2(s string) bool {
	stack := []byte{}

	check := func(ch byte) bool {
		if len(stack) <= 0 {
			return false
		}
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur == '(' && ch == ')' {
			return true
		} else if cur == '{' && ch == '}' {
			return true
		} else if cur == '[' && ch == ']' {
			return true
		}

		return false
	}

	for i := range s {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		case ')', '}', ']':
			if !check(s[i]) {
				return false
			}
		}
	}

	return len(stack) == 0
}
