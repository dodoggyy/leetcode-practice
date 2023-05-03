package easy

// Use hashmap:
// Time Complexity: O(m+n)
// Space Complexity:O(m+n)
// Runtime: 34 ms, faster than 39.22%
// Memory Usage: 7.5 MB, less than 43.14%
func findDifference(nums1 []int, nums2 []int) [][]int {
	hashmap1 := map[int]bool{}
	hashmap2 := map[int]bool{}

	for _, v := range nums1 {
		hashmap1[v] = true
	}

	for _, v := range nums2 {
		hashmap2[v] = true
	}

	tmp1, tmp2 := []int{}, []int{}

	for k := range hashmap1 {
		if !hashmap2[k] {
			tmp1 = append(tmp1, k)
		}
	}

	for k := range hashmap2 {
		if !hashmap1[k] {
			tmp2 = append(tmp2, k)
		}
	}

	return [][]int{tmp1, tmp2}
}
