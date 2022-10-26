package medium

// Use prefix sum + hashmap:
// Time Complexity: O(n)
// Space Complexity:O(min(n,k))
// Runtime: 48 ms, faster than 46.22%
// Memory Usage: 6.8 MB, less than 86.64%
func checkSubarraySum(nums []int, k int) bool {
	size := len(nums)
	if size < 2 {
		return false
	}
	// Need add {0:-1}
	// Due to map cannot match value (idx+1 >= 2) while sub array exist %k = 0 case
	hashmap := map[int]int{0: -1}

	remain := 0
	for i := range nums {
		remain = (remain + nums[i]) % k
		if idx, ok := hashmap[remain]; ok {
			if i-idx >= 2 {
				return true
			}
		} else {
			hashmap[remain] = i
		}
	}

	return false
}
