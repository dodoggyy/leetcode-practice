package easy

// Use Boyer-Moore Voting Algorithm:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 16 ms, faster than 96.03%
// Memory Usage: 6 MB, less than 100.00%
func majorityElement(nums []int) int {
	result := nums[0]
	count := 0
	threshold := len(nums)/2 + 1

	for _, v := range nums {
		if result == v {
			count++
		} else {
			count--
		}
		if count < 0 {
			result = v
			count = 1
		}
		if count == threshold {
			return result
		}
	}

	return result
}
