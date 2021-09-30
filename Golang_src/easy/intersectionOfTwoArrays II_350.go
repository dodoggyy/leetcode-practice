package easy

import "sort"

// Use sort + two pointers:
// Time Complexity: O(nlogn + mlogm)
// Space Complexity:O(min(m,n))
// m: length of nums1
// n: length of nums2
// Runtime: 4 ms, faster than 83.24%
// Memory Usage: 2.8 MB, less than 97.94%
func intersect(nums1 []int, nums2 []int) []int {
	// O(mlogm)
	sort.Ints(nums1)
	// O(nlogn)
	sort.Ints(nums2)

	result := []int{}

	idx1, idx2 := 0, 0
	size1, size2 := len(nums1), len(nums2)

	// O(min(m,n))
	for idx1 < size1 && idx2 < size2 {
		if nums1[idx1] < nums2[idx2] {
			idx1++
		} else if nums1[idx1] > nums2[idx2] {
			idx2++
		} else {
			result = append(result, nums1[idx1])
			idx1++
			idx2++
		}
	}

	return result
}

// Use hashmap:
// Time Complexity: O(m+n)
// Space Complexity:O(min(m,n))
// m: length of nums1
// n: length of nums2
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 3 MB, less than 66.47%
func intersect2(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect2(nums2, nums1)
	}

	hashmap := map[int]int{}
	result := []int{}

	// O(m)
	for _, v := range nums1 {
		hashmap[v]++
	}
	// O(n)
	for _, v := range nums2 {
		if hashmap[v] > 0 {
			result = append(result, v)
			hashmap[v]--
		}

	}

	return result
}

// Use Two pointer to solve in-place while arrays sorted:
// Time Complexity: O(m+n)
// Space Complexity:O(1)
// m: length of nums1
// n: length of nums2
// Runtime: 4 ms, faster than 71.68%
// Memory Usage: 2.8 MB, less than 100.00%
func intersect3(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	size1, size2 := len(nums1), len(nums2)

	idx1, idx2 := 0, 0
	cur := 0

	for idx1 < size1 && idx2 < size2 {
		if nums1[idx1] == nums2[idx2] {
			nums1[cur] = nums1[idx1]
			cur++
			idx1++
			idx2++
		} else if nums1[idx1] < nums2[idx2] {
			idx1++
		} else {
			idx2++
		}
	}

	return nums1[:cur]
}
