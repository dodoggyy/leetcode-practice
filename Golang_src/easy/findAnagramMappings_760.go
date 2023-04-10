package easy

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.3 MB, less than 50%
func anagramMappings(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))

	hashmap := map[int]int{}

	for i, v := range nums2 {
		hashmap[v] = i
	}

	for i, v := range nums1 {
		result[i] = hashmap[v]
	}

	return result
}
