package medium

import (
	"fmt"
	"strconv"
)

// Use hashset + bitwise operation:
// Time Complexity: O(n^2)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.29 MB, less than 39.47%
func findDifferentBinaryString(nums []string) string {
	hashset := map[int]bool{}
	for _, num := range nums {
		v, _ := strconv.ParseInt(num, 2, 32)
		hashset[int(v)] = true
	}

	cnt := 0
	for hashset[cnt] {
		cnt++
	}
	return fmt.Sprintf("%0*b", len(nums[0]), cnt)
}
