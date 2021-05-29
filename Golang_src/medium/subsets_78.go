package medium

import "math"

// Use backtracking:
// Time Complexity: O(n*2^n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.4 MB, less than 16.25%
func subsets(nums []int) [][]int {
	size := len(nums)
	result := [][]int{}
	if size == 0 {
		return result
	}

	var dfs func(int, int, []int)
	dfs = func(idxCur, subsetSize int, subset []int) {
		if subsetSize == 0 {
			tmp := make([]int, len(subset))
			copy(tmp, subset)
			result = append(result, tmp)
			return
		}

		for i := idxCur; i < size; i++ {
			subset = append(subset, nums[i])
			dfs(i+1, subsetSize-1, subset)
			subset = subset[:len(subset)-1]
		}
	}

	for i := 0; i <= size; i++ {
		dfs(0, i, []int{})
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
