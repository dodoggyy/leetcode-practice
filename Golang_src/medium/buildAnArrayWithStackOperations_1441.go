package medium

// Use simulation:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.46 MB, less than 47.83%
func buildArray(target []int, n int) []string {
	result := []string{}

	idx := 0
	cur := 1
	for idx < len(target) {
		if cur == target[idx] {
			result = append(result, "Push")
			idx++
			cur++
		} else {
			for cur < target[idx] {
				result = append(result, "Push", "Pop")
				cur++
			}
		}
	}

	return result
}
