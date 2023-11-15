package medium

// Use greedy:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 73 ms, faster than 75.00%
// Memory Usage: 12.19 MB, less than 25.00%
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	size := len(arr)
	cnt := make([]int, size+1)
	for _, v := range arr {
		cnt[min(v, size)]++
	}

	miss := 0
	for _, v := range cnt[1:] {
		if v == 0 {
			miss++
		} else {
			miss -= min(v-1, miss)
		}
	}

	return size - miss
}
