package medium

import "math"

// Use brute force:
// Time Complexity: O(n^2)
// Space Complexity:O(1)
// Runtime: 80 ms, faster than 10.82%
// Memory Usage: 3.9 MB, less than 45.90%
func minSubArrayLen(target int, nums []int) int {
	result := math.MaxInt32

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				result = min(result, j-i+1)
				break
			}
		}
	}

	if result == math.MaxInt32 {
		return 0
	}

	return result
}

// Use sliding window:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 97.76%
// Memory Usage: 3.8 MB, less than 54.10%
func minSubArrayLen2(target int, nums []int) int {
	result := math.MaxInt32
	size := len(nums)

	l, r := 0, 0
	sum := 0
	for r < size {
		sum += nums[r]
		for sum >= target {
			result = min(result, r-l+1)
			sum -= nums[l]
			l++
		}
		r++
	}

	if result == math.MaxInt32 {
		return 0
	}

	return result
}
