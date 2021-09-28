package medium

// Use DP:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 64 ms, faster than 86.00%
// Memory Usage: 2.2 MB, less than 19.44%
func wiggleMaxLength(nums []int) int {
	size := len(nums)

	up, down := make([]int, size), make([]int, size)

	up[0], down[0] = 1, 1

	for i := 1; i < size; i++ {
		if nums[i] > nums[i-1] {
			up[i] = max(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			up[i] = up[i-1]
			down[i] = max(down[i-1], up[i-1]+1)
		} else {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}

	return max(up[size-1], down[size-1])
}

// Use optimize DP:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 64 ms, faster than 86.00%
// Memory Usage: 2.1 MB, less than 50.00%
func wiggleMaxLength2(nums []int) int {
	size := len(nums)

	up, down := 1, 1

	for i := 1; i < size; i++ {
		if nums[i] > nums[i-1] {
			up = max(up, down+1)
		} else if nums[i] < nums[i-1] {
			down = max(down, up+1)
		}
	}

	return max(up, down)
}
