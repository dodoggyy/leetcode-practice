package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 5 ms, faster than 71.94%
// Memory Usage: 3.42 MB, less than 35.71%
func sumOfSquares(nums []int) int {
	result := 0
	size := len(nums)

	sqrt := func(n int) int {
		return n * n
	}

	for i, v := range nums {
		if size%(i+1) == 0 {
			result += sqrt(v)
		}
	}

	return result
}
