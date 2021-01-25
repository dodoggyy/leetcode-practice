package easy

// Use two pointers:
// Time Complexity: O(m+n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.3 MB, less than 73.11%
func merge(nums1 []int, m int, nums2 []int, n int) {
	index1, index2 := m-1, n-1
	indexCur := m + n - 1

	for index1 >= 0 && index2 >= 0 {
		if nums1[index1] > nums2[index2] {
			nums1[indexCur] = nums1[index1]
			index1--
		} else {
			nums1[indexCur] = nums2[index2]
			index2--
		}
		indexCur--
	}

	for index2 >= 0 {
		nums1[indexCur] = nums2[index2]
		index2--
		indexCur--
	}
}
