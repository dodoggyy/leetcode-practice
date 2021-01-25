package easy

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 56 ms, faster than 82.14%
// Memory Usage: 6.7 MB, less than 71.43%
func findLHS(nums []int) int {
	result := 0
	hashmap := make(map[int]int)

	for _, v := range nums {
		hashmap[v]++
	}

	for k, v := range hashmap {
		if hashmap[k+1] > 0 {
			result = max(result, v+hashmap[k+1])
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Use hashmap for one loop:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 64 ms, faster than 39.29%
// Memory Usage: 6.7 MB, less than 71.43%
func findLHS2(nums []int) int {
	result := 0
	hashmap := make(map[int]int)

	for _, v := range nums {
		hashmap[v]++
		if hashmap[v+1] > 0 {
			result = max(result, hashmap[v]+hashmap[v+1])
		}
		if hashmap[v-1] > 0 {
			result = max(result, hashmap[v]+hashmap[v-1])
		}
	}

	return result
}
