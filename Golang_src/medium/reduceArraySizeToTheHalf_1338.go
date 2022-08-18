package medium

import "sort"

// Use greedy:
// Time Complexity: O(nlogn)
// Space Complexity:O(n)
// Runtime: 279 ms, faster than 5.72%
// Memory Usage: 17 MB, less than 5.72%
func minSetSize(arr []int) int {
	size := len(arr)

	hashmap := map[int]int{}
	for _, v := range arr {
		hashmap[v]++
	}

	keys := make([]int, len(hashmap))
	for k := range hashmap {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return hashmap[keys[i]] > hashmap[keys[j]]
	})

	result, curSize := 0, size
	for _, v := range keys {
		result++
		curSize -= hashmap[v]
		if curSize <= size/2 {
			break
		}
	}

	return result
}
