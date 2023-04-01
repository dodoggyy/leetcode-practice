package easy

// Use iterative:
// Time Complexity: O(n+m)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 36.73%
// Memory Usage: 4.3 MB, less than 33.16%
func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	result := [][]int{}

	idx1, idx2 := 0, 0
	size1, size2 := len(nums1), len(nums2)
	for idx1 < size1 && idx2 < size2 {
		n1, n2 := nums1[idx1], nums2[idx2]
		if n1[0] < n2[0] {
			result = append(result, n1)
			idx1++
		} else if n1[0] == n2[0] {
			result = append(result, []int{n1[0], n1[1] + n2[1]})
			idx1++
			idx2++
		} else {
			result = append(result, n2)
			idx2++
		}
	}

	if idx1 < size1 {
		result = append(result, nums1[idx1:]...)
	}

	if idx2 < size2 {
		result = append(result, nums2[idx2:]...)
	}

	return result
}
