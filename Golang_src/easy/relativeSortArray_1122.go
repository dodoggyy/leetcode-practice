package easy

import "sort"

// Use iterative:
// hash rank: O(n), sort stack: O(logm)
// Time Complexity: O(mlogm + n)
// Space Complexity:O(logm + n)
// Runtime: 3 ms, faster than 48.98%
// Memory Usage: 2.6 MB, less than 20.41%
func relativeSortArray(arr1 []int, arr2 []int) []int {
	hashmap := map[int]int{}
	result := []int{}
	tmp := []int{}

	for _, v := range arr1 {
		hashmap[v]++
	}

	for _, v := range arr2 {
		if val, ok := hashmap[v]; ok {
			for i := 0; i < val; i++ {
				result = append(result, v)
			}
			delete(hashmap, v)
		}
	}

	for k, v := range hashmap {
		for i := 0; i < v; i++ {
			tmp = append(tmp, k)
		}
	}

	sort.Ints(tmp)
	result = append(result, tmp...)

	return result
}

// Use counting sort:
// Time Complexity: O(m + n + upper)
// Space Complexity:O(upper)
// Runtime: 2 ms, faster than 73.47%
// Memory Usage: 2.4 MB, less than 30.61%
func relativeSortArray2(arr1 []int, arr2 []int) []int {
	result := []int{}

	upper := 0
	for _, v := range arr1 {
		if v > upper {
			upper = v
		}
	}
	cnt := make([]int, upper+1)
	for _, v := range arr1 {
		cnt[v]++
	}

	for _, v := range arr2 {
		for ; cnt[v] > 0; cnt[v]-- {
			result = append(result, v)
		}
	}

	for i, v := range cnt {
		for ; v > 0; v-- {
			result = append(result, i)
		}
	}

	return result
}
