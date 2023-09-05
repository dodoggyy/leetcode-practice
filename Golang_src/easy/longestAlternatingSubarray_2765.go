package easy

// Use two pointers:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 12 ms, faster than 61.19%
// Memory Usage: 4.32 MB, less than 41.79%
func alternatingSubarray(nums []int) int {
	result := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(nums); i++ {
		revert := 1
		j := i
		for j+1 < len(nums) && nums[j+1]-nums[j] == revert {
			revert *= -1
			j++
		}
		if i != j {
			result = max(result, j-i+1)
			i = j - 1
		}

	}

	if result == 0 {
		return -1
	}

	return result
}
