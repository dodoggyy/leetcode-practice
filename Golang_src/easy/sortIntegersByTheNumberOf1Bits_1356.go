package easy

import "sort"

// Use bitwise operation and sorting:
// Time Complexity: O(nlogn)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 48.84%
// Memory Usage: 3.41 MB, less than 100.00%
func sortByBits(arr []int) []int {
	cntBit := func(num int) int {
		cnt := 0

		for num > 0 {
			cnt++
			num &= num - 1
		}

		return cnt
	}

	sort.Slice(arr, func(i, j int) bool {
		cntI, cntJ := cntBit(arr[i]), cntBit(arr[j])
		if cntI == cntJ {
			return arr[i] < arr[j]
		}
		return cntI < cntJ
	})

	return arr
}
