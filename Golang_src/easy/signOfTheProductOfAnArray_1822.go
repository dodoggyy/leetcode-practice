package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 7 ms, faster than 50.47%
// Memory Usage: 3.3 MB, less than 98.13%
func arraySign(nums []int) int {
	result := 1

	for _, v := range nums {

		switch {
		case v == 0:
			return 0
		case v < 0:
			result = -result
		}
	}

	return result
}
