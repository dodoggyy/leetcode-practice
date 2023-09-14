package medium

import "sort"

// Use sort + two pointers:
// Time Complexity: O(nlogn + n)
// Space Complexity:O(1)
// Runtime: 109 ms, faster than 73.78%
// Memory Usage: 8.26 MB, less than 83.53%
func maxOperations(nums []int, k int) int {
	result := 0

	sort.Ints(nums)

	l, r := 0, len(nums)-1
	for l < r {
		cur := nums[l] + nums[r]
		if cur == k {
			result++
			l++
			r--
		} else if cur < k {
			l++
		} else {
			r--
		}
	}

	return result
}

// Use sort + two pointers:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 93 ms, faster than 96.52%
// Memory Usage: 9.31 MB, less than 38.05%
func maxOperations2(nums []int, k int) int {
	result := 0

	hashmap := map[int]int{}

	for _, v := range nums {
		if val, ok := hashmap[k-v]; ok && val > 0 {
			result++
			hashmap[k-v]--
		} else {
			hashmap[v]++
		}

	}

	return result
}
