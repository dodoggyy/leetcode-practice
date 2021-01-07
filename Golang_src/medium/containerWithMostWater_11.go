package medium

// Use two pointers:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 24 ms, faster than 31.52%
// Memory Usage: 6.1 MB, less than 7.25%
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	result := 0

	for l < r {
		if height[l] > height[r] {
			result = max(result, height[r]*(r-l))
			r--
		} else {
			result = max(result, height[l]*(r-l))
			l++
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
