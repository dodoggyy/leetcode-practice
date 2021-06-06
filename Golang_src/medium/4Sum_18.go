package medium

import "sort"

// Use sort + two pointers:
// Time Complexity: O(nlogn + n^3) = O(n^3)
// Space Complexity:O(n)
// Runtime: 4 ms, faster than 96.43%
// Memory Usage: 2.8 MB, less than 61.16%
func fourSum(nums []int, target int) [][]int {
	result := [][]int{}

	// O(nlogn)
	sort.Ints(nums)

	size := len(nums)

	// O(n^3)
	for i := 0; i < size-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		// ignore same case or less than target case
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[size-3]+nums[size-2]+nums[size-1] < target {
			continue
		}

		for j := i + 1; j < size-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			// ignore same case or less than target case
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[size-2]+nums[size-1] < target {
				continue
			}

			left, right := j+1, size-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
					for left < right && nums[left+1] == nums[left] {
						left++
					}
					for left < right && nums[right-1] == nums[right] {
						right--
					}
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
		}
	}

	return result
}
