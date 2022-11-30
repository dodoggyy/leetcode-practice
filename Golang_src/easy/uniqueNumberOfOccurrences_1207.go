package easy

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 72.55%
// Memory Usage: 2.3 MB, less than 100.00%
func uniqueOccurrences(arr []int) bool {
	hashmap := map[int]int{}
	for _, v := range arr {
		hashmap[v]++
	}

	hashset := map[int]bool{}

	for _, v := range hashmap {
		if hashset[v] {
			return false
		}
		hashset[v] = true
	}

	return true
}
