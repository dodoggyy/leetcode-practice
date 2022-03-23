package medium

// Use Work Backwards:
// Time Complexity: O(log(target))
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.9 MB, less than 28.57%
func brokenCalc(startValue int, target int) int {
	result := 0

	for startValue < target {
		result++
		if target%2 == 1 {
			target++
		} else {
			target /= 2
		}
	}

	return result + startValue - target
}
