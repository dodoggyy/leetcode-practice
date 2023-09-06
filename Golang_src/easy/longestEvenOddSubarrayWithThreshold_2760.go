package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 29 ms, faster than 100.00%
// Memory Usage: 7.37 MB, less than 67.33%
func longestAlternatingSubarray(nums []int, threshold int) int {
	result := 0
	cnt := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := range nums {
		if nums[i] > threshold {
			cnt = 0
		} else {
			if i == 0 || nums[i]%2 == nums[i-1]%2 {
				cnt = 1
			} else {
				cnt++
			}
			// if head is odd, need to ignore it
			if nums[i-cnt+1]%2 == 1 {
				result = max(result, cnt-1)
			} else {
				result = max(result, cnt)
			}
		}

	}

	return result
}
