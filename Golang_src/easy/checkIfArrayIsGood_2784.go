package easy

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 3.11 MB, less than 21.69%
func isGood(nums []int) bool {
	hashmap := map[int]int{}
	numTwice := len(nums) - 1

	for _, v := range nums {
		hashmap[v]++
		if hashmap[v] > 1 && v != numTwice {
			return false
		}
		if v > numTwice {
			return false
		}
	}

	if hashmap[numTwice] > 2 {
		return false
	}

	return true
}
