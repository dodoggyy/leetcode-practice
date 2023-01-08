package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 172  ms, faster than 79.8%
// Memory Usage: 9 MB, less than 46.43%
func findMaxAverage(nums []int, k int) float64 {
	size := len(nums)

	tmp := float64(0)

	for i := 0; i < k; i++ {
		tmp += float64(nums[i])
	}
	result := tmp

	for i := k; i < size; i++ {
		tmp += float64(nums[i] - nums[i-k])
		if result < tmp {
			result = tmp
		}
	}

	return result / float64(k)
}
