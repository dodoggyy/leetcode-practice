package easy

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 1 ms, faster than 78.10%
// Memory Usage: 2.02 MB, less than 19.68%
func numIdenticalPairs(nums []int) int {
	result := 0

	hashmap := map[int]int{}
	for _, v := range nums {
		hashmap[v]++
	}

	// c(n,2)
	cal := func(times int) int {
		if times < 2 {
			return 0
		}
		return times * (times - 1) / 2
	}

	for _, v := range hashmap {
		result += cal(v)
	}

	return result
}
