package medium

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 155 ms, faster than 14.29%
// Memory Usage: 10.2 MB, less than 14.29%
func zeroFilledSubarray(nums []int) int64 {
	result := int64(0)
	cnt := 0
	for _, v := range nums {
		if v == 0 {
			cnt++
			result += int64(cnt)
		} else {
			cnt = 0
		}
	}

	return result
}
