package easy

// Use HashSet:
// Time Complexity: O(m+n)
// Space Complexity:O(m+n)
// m : length of nums1
// n : length of nums2
// Runtime: 4 ms, faster than 65.14%
// Memory Usage: 3.2 MB, less than 45.41%
func intersection(nums1 []int, nums2 []int) []int {
	hashmap := map[int]bool{}

	for _, v := range nums1 {
		hashmap[v] = true
	}

	result := []int{}

	for _, v := range nums2 {
		if _, ok := hashmap[v]; ok {
			result = append(result, v)
			delete(hashmap, v)
		}
	}

	return result
}
