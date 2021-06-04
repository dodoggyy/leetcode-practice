package medium

// Use partition + hash table:
// Time Complexity: O(n^2)
// Space Complexity:O(n^2)
// Runtime: 120 ms, faster than 60.19%
// Memory Usage: 6.8 MB, less than 47.22%
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	hashmap := map[int]int{}
	result := 0

	// O(n^2)
	for _, v := range nums1 {
		for _, w := range nums2 {
			hashmap[v+w]++
		}
	}

	// O(n^2)
	for _, v := range nums3 {
		for _, w := range nums4 {
			result += hashmap[-v-w]
		}
	}

	return result
}
