package medium

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 99 ms, faster than 92.00%
// Memory Usage: 9.48 MB, less than 72.00%
func minOperations2(nums []int) int {
	result := 0

	hashmap := map[int]int{}
	for _, v := range nums {
		hashmap[v]++
	}

	for _, v := range hashmap {
		if v == 1 {
			return -1
		}
		if v%3 == 0 {
			result += v / 3
		} else if v%3 == 1 {
			result += (v-4)/3 + 2
		} else {
			result += (v-2)/3 + 1
		}
	}

	return result
}
