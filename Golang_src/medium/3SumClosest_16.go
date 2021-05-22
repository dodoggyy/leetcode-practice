package medium

import (
	"math"
	"sort"
)

// Use sort + two pointers:
// Time Complexity: O(nlogn+ n^2) = O(n^2)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 42.60%
// Memory Usage: 3 MB, less than 6.73%
func threeSumClosest(nums []int, target int) int {
	// O(nlogn)
	sort.Ints(nums)
	result := math.MaxInt32
	size := len(nums)

	// Total: O(n^2)
	// O(n)
	for i := 0; i < size-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, size-1

		// O(n)
		for left < right {
			cur := nums[left] + nums[right] + nums[i]
			if cur == target {
				return cur
			}

			if absolute(cur-target) < absolute(result-target) {
				result = cur
			}

			if cur > target {
				for left < right && nums[right-1] == nums[right] {
					right--
				}
				right--
			} else {
				for left < right && nums[left+1] == nums[left] {
					left++
				}
				left++
			}
		}

	}

	return result
}

func absolute(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}
