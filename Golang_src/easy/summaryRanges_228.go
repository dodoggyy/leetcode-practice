package easy

import "strconv"

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 1 ms, faster than 69.60%
// Memory Usage: 2 MB, less than 41.39%
func summaryRanges(nums []int) []string {
	result := []string{}

	if len(nums) == 0 {
		return result
	}

	start, end := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if end+1 == nums[i] {
			end = nums[i]
		} else {
			if start != end {
				result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(end))
			} else {
				result = append(result, strconv.Itoa(start))
			}

			start, end = nums[i], nums[i]
		}
	}

	if start != end {
		result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(end))
	} else {
		result = append(result, strconv.Itoa(start))
	}

	return result
}
