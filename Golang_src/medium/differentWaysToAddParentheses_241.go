package medium

import "strconv"

// Use divide and conquer
// Time Complexity: O(n^3)
// Space Complexity:O(mn)
// m for operator count
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.3 MB, less than 19.44%
func diffWaysToCompute(expression string) []int {
	if isDigit(expression) {
		tmp, _ := strconv.Atoi(expression)
		return []int{tmp}
	}

	result := []int{}

	for idx, ch := range expression {
		if ch == '+' || ch == '-' || ch == '*' {
			left := diffWaysToCompute(expression[:idx])
			right := diffWaysToCompute(expression[idx+1:])

			for _, valLeft := range left {
				for _, valRight := range right {
					var tmpResult int
					switch ch {
					case '+':
						tmpResult = valLeft + valRight
					case '-':
						tmpResult = valLeft - valRight
					case '*':
						tmpResult = valLeft * valRight
					}
					result = append(result, tmpResult)
				}
			}
		}
	}

	return result
}

func isDigit(str string) bool {
	_, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return true
}
