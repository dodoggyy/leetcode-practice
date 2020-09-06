package easy

// Use hashmap one pass
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 3.8 MB, less than 51.15%
func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for i, value := range nums {
		complement := target - value
		if j, exist := hashmap[complement]; exist {
			return []int{j, i}
		}
		hashmap[value] = i
	}

	return []int{}
}
