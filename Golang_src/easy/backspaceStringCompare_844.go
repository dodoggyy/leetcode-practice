package easy

// Use stack:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.9 MB, less than 100.00%
func backspaceCompare(s string, t string) bool {
	preprocess := func(str string) []byte {
		stack := []byte{}
		for i := range str {
			if str[i] == '#' {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			} else {
				stack = append(stack, str[i])
			}
		}

		return stack
	}

	stack1, stack2 := preprocess(s), preprocess(t)
	if len(stack1) != len(stack2) {
		return false
	}
	for i := range stack1 {
		if stack1[i] != stack2[i] {
			return false
		}
	}

	return true
}

// Use two pointers:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 2 ms, faster than 55.59%
// Memory Usage: 1.9 MB, less than 100.00%
func backspaceCompare2(s string, t string) bool {
	cur1, cur2 := len(s)-1, len(t)-1
	skip1, skip2 := 0, 0

	for cur1 >= 0 || cur2 >= 0 {
		for cur1 >= 0 {
			if s[cur1] == '#' {
				skip1++
				cur1--
			} else if skip1 > 0 {
				skip1--
				cur1--
			} else {
				break
			}
		}
		for cur2 >= 0 {
			if t[cur2] == '#' {
				skip2++
				cur2--
			} else if skip2 > 0 {
				skip2--
				cur2--
			} else {
				break
			}
		}
		if cur1 < 0 || cur2 < 0 {
			break
		}
		if s[cur1] != t[cur2] {
			return false
		}
		cur1--
		cur2--
	}

	return cur1 == cur2
}
