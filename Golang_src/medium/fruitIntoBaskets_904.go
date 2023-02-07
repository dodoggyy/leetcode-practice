package medium

// Use sliding windows:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 100 ms, faster than 97.83%
// Memory Usage: 8.5 MB, less than 78.26%
func totalFruit(fruits []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	result := 0
	cnt := map[int]int{}
	left := 0
	for right, v := range fruits {
		cnt[v]++
		for len(cnt) > 2 {
			val := fruits[left]
			cnt[val]--
			if cnt[val] == 0 {
				delete(cnt, val)
			}
			left++
		}
		result = max(result, right-left+1)
	}

	return result
}
