package medium

// Use bit operation:
// Time Complexity: O(nlogc)
// Space Complexity:O(1)
// Runtime: 3 ms, faster than 46.22%
// Memory Usage: 6.8 MB, less than 86.64%
func singleNumberII(nums []int) int {
	result := int32(0)

	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, v := range nums {
			total += (int32(v) >> i) & 1
		}
		if total%3 > 0 {
			result |= 1 << i
		}
	}

	return int(result)
}
