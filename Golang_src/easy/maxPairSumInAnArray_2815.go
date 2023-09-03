package easy

import "math"

// Use one pass iterative:
// n: lengths of nums
// U: max(nums)
// Time Complexity: O(nlogU)
// Space Complexity:O(1)
// Runtime: 16 ms, faster than 63.79%
// Memory Usage: 5.03 MB, less than 95.40%
func maxSum(nums []int) int {
	result := -1
	records := [10]int{}
	for i := range records {
		records[i] = math.MinInt32
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := range nums {
		maxDigit := 0
		num := nums[i]
		for num != 0 {
			maxDigit = max(maxDigit, num%10)
			num /= 10
		}

		if records[maxDigit] != math.MinInt32 {
			result = max(result, records[maxDigit]+nums[i])
		}

		records[maxDigit] = max(records[maxDigit], nums[i])
	}

	return result
}
