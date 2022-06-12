package medium

// Use sliding window & hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 212 ms, faster than 66.67%
// Memory Usage: 9.1 MB, less than 72.00%
func maximumUniqueSubarray(nums []int) int {
	l, r := 0, 0
	result := 0
	hashmap := map[int]int{}
	tmp := 0

	for r < len(nums) {
		if _, ok := hashmap[nums[r]]; ok {
			for nums[l] != nums[r] && l <= r {
				tmp -= nums[l]
				delete(hashmap, nums[l])
				l++

			}
			tmp -= nums[l]
			delete(hashmap, nums[l])
			l++
		}
		tmp += nums[r]
		hashmap[nums[r]] = r
		r++

		result = max(result, tmp)
	}

	return result
}
