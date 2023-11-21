package medium

import "sort"

// Use binary search:
// Time Complexity: O(target*n)
// Space Complexity:O(target)
// Runtime: 174 ms, faster than 72.63%
// Memory Usage: 8.86 MB, less than 94.74%
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result := [][]int{}
	m, n := len(nums1), len(nums2)

	l, r := nums1[0]+nums2[0], nums1[m-1]+nums2[n-1]
	sumMid := l + sort.Search(r-l, func(sum int) bool {
		sum += l
		cnt := 0
		i, j := 0, n-1
		for i < m && j >= 0 {
			if nums1[i]+nums2[j] > sum {
				j--
			} else {
				cnt += j + 1
				i++
			}
		}
		return cnt >= k
	})

	// find pairs which sum less than sumMid
	i := n - 1
	for _, v1 := range nums1 {
		for i >= 0 && v1+nums2[i] >= sumMid {
			i--
		}
		for _, v2 := range nums2[:i+1] {
			result = append(result, []int{v1, v2})
			if len(result) == k {
				return result
			}
		}
	}

	// find pairs which sum equal to sumMid
	i = n - 1
	for _, v1 := range nums1 {
		for i >= 0 && v1+nums2[i] > sumMid {
			i--
		}
		for j := i; j >= 0 && v1+nums2[j] == sumMid; j-- {
			result = append(result, []int{v1, nums2[j]})
			if len(result) == k {
				return result
			}
		}
	}

	return result
}
