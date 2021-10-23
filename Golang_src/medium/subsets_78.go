package medium

import "math"

// Use backtracking:
// Time Complexity: O(n*2^n)
// Space Complexity:O(n)
// Runtime: 2 ms, faster than 27.30%
// Memory Usage: 2.6 MB, less than 20.73%
func subsets(nums []int) [][]int {
	result := [][]int{}
	keep := []int{}
	size := len(nums)

	var dfs func(idx, subSize int)
	dfs = func(idx, subSize int) {
		if subSize == 0 {
			tmp := make([]int, len(keep))
			copy(tmp, keep)
			result = append(result, tmp)
			return
		}

		for i := idx; i < size; i++ {
			keep = append(keep, nums[i])
			dfs(i+1, subSize-1)
			keep = keep[:len(keep)-1]
		}
	}

	for i := 0; i <= size; i++ {
		dfs(0, i)
	}

	return result
}

// Use bitmask:
// Time Complexity: O(n*2^n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.4 MB, less than 16.25%
func subsets2(nums []int) [][]int {
	size := len(nums)
	result := [][]int{}
	if size == 0 {
		return result
	}

	last := int(math.Pow(2.0, float64(size)))

	for i := 0; i < last; i++ {
		tmp := []int{}
		for j := 0; j < size; j++ {
			if (i & (1 << j)) != 0 {
				tmp = append(tmp, nums[j])
			}
		}
		result = append(result, tmp)
	}

	return result
}
