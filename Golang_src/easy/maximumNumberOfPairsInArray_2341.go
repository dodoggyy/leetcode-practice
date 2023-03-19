package easy

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100%
// Memory Usage: 2.4 MB, less than 23.81%
func numberOfPairs(nums []int) []int {
	hashmap := map[int]int{}
	for _, v := range nums {
		hashmap[v]++
	}

	result := make([]int, 2)

	for _, v := range hashmap {
		result[0] += v / 2
		result[1] += v % 2
	}

	return result
}
