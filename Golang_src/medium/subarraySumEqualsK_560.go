package medium

// Use brute force:
// Time Complexity: O(n^2)
// Space Complexity:O(1)
// Runtime: 888 ms, faster than 23.43%
// Memory Usage: 6.2 MB, less than 81.12%
func subarraySum(nums []int, k int) int {
	result := 0

	size := len(nums)
	for end := 0; end < size; end++ {
		sum := 0
		for start := end; start >= 0; start-- {
			sum += nums[start]
			if sum == k {
				result++
			}
		}
	}

	return result
}

// Use prefix sum + hash table:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 44 ms, faster than 79.37%
// Memory Usage: 7.2 MB, less than 21.68%
func subarraySum2(nums []int, k int) int {
	result := 0
	pre := 0
	hashmap := map[int]int{}
	hashmap[0] = 1

	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if val, ok := hashmap[pre-k]; ok {
			result += val
		}
		hashmap[pre]++
	}

	return result
}
