package easy

// Use comparison:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 43.40%
// Memory Usage: 2.9 MB, less than 49.6%
func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	result := 0

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	if k > 0 {
		result += min(numOnes, k) * 1
		k -= numOnes
	}
	if k > 0 {
		result += min(numZeros, k) * 0
		k -= numZeros
	}
	if k > 0 {
		result += min(numNegOnes, k) * -1
	}

	return result
}
