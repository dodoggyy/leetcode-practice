package easy

import "math"

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 34 ms, faster than 30.77%
// Memory Usage: 7.1 MB, less than 23.8%
func findShortestSubArray(nums []int) int {
	result := math.MaxInt32

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	maxDeg := 0

	// [1,2,2,3,1]
	// 1: [0,4]
	// 2: [1,2]
	// 3: [3]
	hashmap := map[int][]int{}
	for i, v := range nums {
		if val, ok := hashmap[v]; ok {
			val = append(val, i)
			hashmap[v] = val
			maxDeg = max(maxDeg, len(val))
		} else {
			hashmap[v] = []int{i}
			maxDeg = max(maxDeg, 1)
		}
	}

	for _, v := range hashmap {
		if len(v) == maxDeg {
			result = min(result, v[len(v)-1]-v[0]+1)
		}
	}

	return result
}
