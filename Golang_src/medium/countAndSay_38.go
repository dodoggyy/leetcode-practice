package medium

import "strconv"

// Use iterative:
// Time Complexity: O(n*k)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 67.01%
// Memory Usage: 7.4 MB, less than 39.69%
func countAndSay(n int) string {
	result := "1"
	// O(n)
	for i := 1; i < n; i++ {
		result = helperCnt(result)
	}

	return result
}

func helperCnt(s string) string {
	result := ""
	cnt := 1
	cur := s[0]
	// O(n*k)
	for i := 1; i < len(s); i++ {
		if s[i] == cur {
			cnt++
		} else {
			result += strconv.Itoa(cnt) + string(cur)
			cnt = 1
			cur = s[i]
		}
	}

	result += strconv.Itoa(cnt) + string(cur)

	return result
}
