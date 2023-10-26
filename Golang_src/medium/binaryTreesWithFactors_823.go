package medium

import "sort"

// Use DP:
// Time Complexity: O(n^2)
// Space Complexity:O(n)
// Runtime: 17 ms, faster than 62.50%
// Memory Usage: 3.93 MB, less than 12.50%
func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)

	hashmap := map[int]int{}
	for i, v := range arr {
		hashmap[v] = i
	}

	f := make([]int, len(arr))

	result := 0

	for i, v := range arr {
		f[i] = 1
		for j, x := range arr[:i] {
			if v%x == 0 {
				if k, ok := hashmap[v/x]; ok {
					f[i] += f[j] * f[k]
				}
			}
		}
		result += f[i]
	}

	return result % 1000000007
}
