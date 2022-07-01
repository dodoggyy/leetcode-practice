package easy

import "sort"

// Use greedy:
// Time Complexity: O(nlogn)
// Space Complexity:O(1)
// Runtime: 28 ms, faster than 69.72%
// Memory Usage: 6.9 MB, less than 23.85%
func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		m, n := boxTypes[i], boxTypes[j]
		return m[1] > n[1]
	})

	result := 0

	for _, v := range boxTypes {
		size := min(truckSize, v[0])
		result += size * v[1]
		truckSize -= size
	}

	return result
}
