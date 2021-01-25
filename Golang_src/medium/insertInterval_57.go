package medium

// Use simulation:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 78.72%
// Memory Usage: 4.8 MB, less than 85.64%
func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	left, right := newInterval[0], newInterval[1]
	merged := false

	for _, interval := range intervals {
		if interval[0] > right { // interval locate newInterval right and no overlap
			if !merged {
				result = append(result, []int{left, right})
				merged = true
			}
			result = append(result, interval)
		} else if interval[1] < left { // interval locate newInterval left and no overlap
			result = append(result, interval)
		} else { // interval overlap
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	if !merged {
		result = append(result, []int{left, right})
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
