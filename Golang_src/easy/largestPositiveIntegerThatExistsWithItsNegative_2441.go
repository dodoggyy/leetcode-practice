package easy

// Use hashset:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 20 ms, faster than 45.71%
// Memory Usage: 6.9 MB, less than 22.86%
func findMaxK(nums []int) int {
	hashset := map[int]struct{}{}
	result := -1

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for _, v := range nums {
		if v < 0 {
			hashset[v] = struct{}{}
		} else {
			if _, ok := hashset[-v]; ok {
				result = max(result, v)
			}
		}
	}

	return result
}
