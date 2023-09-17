package medium

// Use sliding window:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 53 ms, faster than 57.67%
// Memory Usage: 7.82 MB, less than 38.17%
func longestSubarray(nums []int) int {
	result := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	l, r := 0, 0
	cnt := 0
	curZero := -1
	for r < len(nums) {
		if nums[r] == 0 {
			if curZero >= l && curZero < r {
				l = curZero + 1
			}
			curZero = r
		} else {
			cnt++
			if curZero >= l && curZero < r {
				result = max(result, r-l)
			} else {
				result = max(result, r-l+1)
			}

		}
		r++
	}

	if curZero == -1 {
		return result - 1
	}

	return result
}
