package medium

// Use hashmap:
// Time Complexity: O(n^2)
// Space Complexity:O(n^2)
// Runtime: 63 ms, faster than 95.83%
// Memory Usage: 8.40 MB, less than 45.83%
func countNicePairs(nums []int) int {
	result := 0

	revert := func(n int) int {
		tmp := 0
		for n > 0 {
			tmp *= 10
			tmp += n % 10
			n /= 10
		}
		return tmp
	}

	hashmap := map[int]int{}
	for _, v := range nums {
		rev := revert(v)
		result += hashmap[v-rev]
		hashmap[v-rev]++
	}

	return result % (1e9 + 7)
}
