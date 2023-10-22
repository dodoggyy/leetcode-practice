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

// Use sliding window + hashset:
// Time Complexity: O(n)
// Space Complexity:O(k)
// Runtime: 112 ms, faster than 48.02%
// Memory Usage: 7.73 MB, less than 99.04%
func containsNearbyDuplicate2(nums []int, k int) bool {
	hashset := map[int]bool{}
	for i, v := range nums {
		if i > k {
			delete(hashset, nums[i-k-1])
		}
		if hashset[v] {
			return true
		}

		hashset[v] = true
	}
	return false
}
