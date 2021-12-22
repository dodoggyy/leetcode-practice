package easy

import "strconv"

// You are given an inclusive range [lower, upper] and a sorted unique integer array nums, where all elements are in the inclusive range.
// A number x is considered missing if x is in the range [lower, upper] and x is not in nums.
// Return the smallest sorted list of ranges that cover every missing number exactly. That is, no element of nums is in any of the ranges, and each missing number is in one of the ranges.
// Each range [a,b] in the list should be output as:

// "a->b" if a != b
// "a" if a == b

// Example 1:
// Input: nums = [0,1,3,50,75], lower = 0, upper = 99
// Output: ["2","4->49","51->74","76->99"]
// Explanation: The ranges are:
// [2,2] --> "2"
// [4,49] --> "4->49"
// [51,74] --> "51->74"
// [76,99] --> "76->99"

// Example 2:
// Input: nums = [-1], lower = -1, upper = -1
// Output: []
// Explanation: There are no missing ranges since there are no missing numbers.

// Use linear scan:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 100.00%
func findMissingRanges(nums []int, lower int, upper int) []string {
	result := []string{}

	transfer := func(low, upp int) string {
		if low == upp {
			return strconv.Itoa(low)
		}

		return strconv.Itoa(low) + "->" + strconv.Itoa(upp)
	}

	pre := lower - 1
	for i := range nums {
		cur := nums[i]
		if pre+1 <= cur-1 {
			result = append(result, transfer(pre+1, cur-1))
		}
		pre = cur
	}
	if pre+1 <= upper {
		result = append(result, transfer(pre+1, upper))
	}

	return result
}
