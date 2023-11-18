package hard

// Use binary search:
// Time Complexity: O(log(m+n))
// Space Complexity:O(1)
// Runtime: 13 ms, faster than 45.17%
// Memory Usage: 4.79 MB, less than 100.00%
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	getKthElement := func(k int) int {
		idx1, idx2 := 0, 0
		for {
			if idx1 == len(nums1) {
				return nums2[idx2+k-1]
			}
			if idx2 == len(nums2) {
				return nums1[idx1+k-1]
			}
			if k == 1 {
				return min(nums1[idx1], nums2[idx2])
			}
			idxNew1 := min(idx1+k/2, len(nums1)) - 1
			idxNew2 := min(idx2+k/2, len(nums2)) - 1
			pivot1, pivot2 := nums1[idxNew1], nums2[idxNew2]
			if pivot1 <= pivot2 {
				k -= (idxNew1 - idx1 + 1) // k=k−k/2
				idx1 = idxNew1 + 1
			} else {
				k -= (idxNew2 - idx2 + 1)
				idx2 = idxNew2 + 1
			}
		}
	}

	size := len(nums1) + len(nums2)
	if size%2 == 1 {
		idxMid := size / 2
		return float64(getKthElement(idxMid + 1))
	} else {
		idxMid1, idxMid2 := size/2-1, size/2
		return float64((getKthElement(idxMid1+1) + getKthElement(idxMid2+1))) / 2.0
	}
}
