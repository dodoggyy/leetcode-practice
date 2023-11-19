package medium

import "sort"

// Use hashmap + sorting:
// Time Complexity: O(nlogn)
// Space Complexity:O(n)
// Runtime: 250 ms, faster than 100.00%
// Memory Usage: 12.42 MB, less than 100.00%
func reductionOperations(nums []int) int {
	hashmap := map[int]int{}
	for _, v := range nums {
		hashmap[v]++
	}
	arr := []int{}
	for k := range hashmap {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	result := 0
	tmp := 0

	for i := len(arr) - 1; i >= 1; i-- {
		tmp += hashmap[arr[i]]
		result += tmp
	}

	return result
}

// Use sorting:
// Time Complexity: O(nlogn)
// Space Complexity:O(logn)
// Runtime: 169 ms, faster than 100.00%
// Memory Usage: 7.18 MB, less than 100.00%
func reductionOperations2(nums []int) int {
	result := 0
	sort.Ints(nums)
	cnt := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			cnt++
		}
		result += cnt
	}

	return result
}
