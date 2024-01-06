package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 82.39%
// Memory Usage: 2.98 MB, less than 99.30%
func hasTrailingZeros(nums []int) bool {
	hasTrail := false
	for _, v := range nums {
		if v&1 == 0 {
			if hasTrail {
				return true
			} else {
				hasTrail = true
			}
		}
	}

	return false
}
