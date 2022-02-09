package medium

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 7 ms, faster than 95.56%
// Memory Usage: 6.4 MB, less than 33.33%
func findPairs(nums []int, k int) int {
	result := 0
	hashmap := map[int]int{}

	for _, v := range nums {
		hashmap[v]++
	}

	for key := range hashmap {
		if k == 0 {
			if hashmap[key] > 1 {
				result++
			}
		} else {
			if _, ok := hashmap[key-k]; ok {
				result++
			}
		}

	}

	return result
}
