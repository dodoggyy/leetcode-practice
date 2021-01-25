package easy

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 20 ms, faster than 82.98%
// Memory Usage: 6.3 MB, less than 35.32%
func containsDuplicate(nums []int) bool {
	hashmap := make(map[int]bool)

	for _, v := range nums {
		_, exist := hashmap[v]
		if exist {
			return true
		} else {
			hashmap[v] = true
		}
	}

	return false
}
