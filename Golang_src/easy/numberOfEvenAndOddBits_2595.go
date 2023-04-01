package easy

// Use iterative:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 3 ms, faster than 52.80%
// Memory Usage: 2.2 MB, less than 91.93%
func evenOddBit(n int) []int {
	result := make([]int, 2)
	isOdd := true

	for n > 0 {
		if isOdd {
			result[0] += n & 1
		} else {
			result[1] += n & 1
		}

		isOdd = !isOdd
		n >>= 1
	}

	return result
}
