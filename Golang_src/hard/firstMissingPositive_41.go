package hard

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 172 ms, faster than 5.05%
// Memory Usage: 40.6 MB, less than 5.05%
func firstMissingPositive(nums []int) int {
	hashmap := map[int]bool{}

	for _, v := range nums {
		hashmap[v] = true
	}

	for i := 1; i <= len(nums); i++ {
		if !hashmap[i] {
			return i
		}
	}

	return len(nums) + 1
}

// Use replace:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 172 ms, faster than 5.05%
// Memory Usage: 40.6 MB, less than 5.05%
func firstMissingPositive2(nums []int) int {
	size := len(nums)
	hasOne := false

	// O(n)
	for i := range nums {
		if nums[i] == 1 {
			hasOne = true
		}
	}
	// check 1 exist or not to prevent e.g. [4,5,6,7] case
	// if not check, it will perform like [-1,1,1,1] then return 2
	if !hasOne {
		return 1
	}

	// O(n)
	// replace out of index case as 1
	for i := range nums {
		if nums[i] <= 0 || nums[i] > size {
			nums[i] = 1
		}
	}

	// O(n)
	// set index of value as negative to ensure value is exist
	// nums[abs(nums[i] - 1)] = -abs(nums[abs(nums[i]) - 1])
	for i := range nums {
		tmp := abs(nums[i]) - 1
		nums[tmp] = -abs(nums[tmp])
	}

	// O(n)
	// check positvie means not in nums value
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return size + 1
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
