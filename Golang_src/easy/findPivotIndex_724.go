package easy

// Use prefix sum:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 17 ms, faster than 88.20%
// Memory Usage: 6.2 MB, less than 94.81%
func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	sumLeft := 0
	for i := range nums {
		if sumLeft == sum-sumLeft-nums[i] {
			return i
		}
		sumLeft += nums[i]
	}

	return -1
}
