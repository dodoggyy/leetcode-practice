package medium

// Use Bit Manipulation:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 100.00%
// Memory Usage: 4.2 MB, less than 41.18%
func singleNumber(nums []int) []int {
	result := []int{0, 0}

	tmp := 0

	for _, num := range nums {
		tmp ^= num
	}

	tmp &= (-tmp)

	for _, num := range nums {
		if tmp&num != 0 {
			result[0] ^= num
		} else {
			result[1] ^= num
		}
	}

	return result
}
