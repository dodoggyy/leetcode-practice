package medium

// Use backtracking:
// Time Complexity: O(2^n)
// Space Complexity:O(n)
// Runtime: 176 ms, faster than 15.79%
// Memory Usage: 6.5 MB, less than 57.89%
func maxLength(arr []string) int {
	checkUniq := func(str string) bool {
		if len(str) == 0 {
			return true
		}

		hashset := map[byte]bool{}
		for i := range str {
			if _, ok := hashset[str[i]]; ok {
				return false
			}
			hashset[str[i]] = true
		}

		return true
	}

	var dfs func(str string, start, end int) string
	dfs = func(str string, start, end int) string {
		if len(str) > 26 {
			return ""
		}
		if start > end {
			if checkUniq(str) {
				return str
			}
			return ""
		}

		noAddStr := dfs(str, start+1, end)
		addStr := ""

		if checkUniq(str + arr[start]) {
			addStr = dfs(str+arr[start], start+1, end)
		}

		if len(noAddStr) > len(addStr) {
			return noAddStr
		}
		return addStr
	}

	result := dfs("", 0, len(arr)-1)

	return len(result)
}
