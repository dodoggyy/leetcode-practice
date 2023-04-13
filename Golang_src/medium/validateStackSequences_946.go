package medium

// Use stack:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 9 ms, faster than 38.24%
// Memory Usage: 3.8 MB, less than 11.76%
func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	i := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}

	return len(stack) == 0
}

// Use simulation:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 8 ms, faster than 41.18%
// Memory Usage: 3.6 MB, less than 100%
func validateStackSequences2(pushed []int, popped []int) bool {
	i, j := -1, 0
	for _, v := range pushed {
		i++
		pushed[i] = v
		for i > -1 && popped[j] == pushed[i] {
			i--
			j++
		}
	}

	return i == -1
}
