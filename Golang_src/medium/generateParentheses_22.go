package medium

// Use backtracking:
// Time Complexity: O((4^n)/(n^1/2))
// Space Complexity:O(n)
// Runtime: 4 ms, faster than 14.53%
// Memory Usage: 3.1 MB, less than 26.74%
func generateParenthesis(n int) []string {
	result := []string{}

	var dfs func(int, int, string)
	dfs = func(left, right int, str string) {
		if left < 0 || right < 0 {
			return
		}

		if left == 0 && right == 0 {
			result = append(result, str)
			return
		}
		if left == right {
			dfs(left-1, right, str+"(")
		} else if left < right {
			dfs(left, right-1, str+")")
			dfs(left-1, right, str+"(")
		}
	}

	dfs(n, n, "")

	return result
}
