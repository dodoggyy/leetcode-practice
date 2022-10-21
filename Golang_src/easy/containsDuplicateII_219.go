package easy

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 244 ms, faster than 35.81%
// Memory Usage: 14.3 MB, less than 15.88%
func containsNearbyDuplicate(nums []int, k int) bool {
	hashmap := map[int]int{}
	for i := range nums {
		if val, ok := hashmap[nums[i]]; ok {
			if i-val <= k {
				return true
			}
		}
		hashmap[nums[i]] = i
	}

	return false
}
