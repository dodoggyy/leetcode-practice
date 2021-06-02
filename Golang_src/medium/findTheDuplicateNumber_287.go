package medium

// Use hash table:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 0.0 MB, less than 100%
func findDuplicate(nums []int) int {
	hashmap := map[int]bool{}
	for _, v := range nums {
		if hashmap[v] {
			return v
		}
		hashmap[v] = true
	}

	return -1
}

// Use slow/fast pointers:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 184 ms, faster than 5.62%
// Memory Usage: 8.9 MB, less than 22.49%
func findDuplicate2(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}

	fast = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}

	return fast
}
