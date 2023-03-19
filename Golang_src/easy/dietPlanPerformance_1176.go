package easy

// Use sliding window:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 27 ms, faster than 33.33%
// Memory Usage: 8.2 MB, less than 33.33%
func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	result := 0

	l, r := 0, 0
	cur := 0
	for r < len(calories) {
		cur += calories[r]

		if r-l+1 > k {
			cur -= calories[l]
			l++
		}

		if r-l+1 == k {
			if cur > upper {
				result++
			} else if cur < lower {
				result--
			}
		}

		r++
	}

	return result
}
