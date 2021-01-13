package medium

import "sort"

// Use sort + two pointers:
// Time Complexity: O(n^2)
// Space Complexity:O(logn)
// Runtime: 48 ms, faster than 46.22%
// Memory Usage: 6.8 MB, less than 86.64%
func threeSum(nums []int) [][]int {
	size := len(nums)
	result := make([][]int, 0)

	sort.Ints(nums)

	for i := 0; i < size-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		complement := -1 * nums[i]
		left, right := i+1, size-1
		for left < right {
			tmp := nums[left] + nums[right]
			if tmp == complement {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if tmp > complement {
				right--
			} else {
				left++
			}
		}
	}

	return result
}
