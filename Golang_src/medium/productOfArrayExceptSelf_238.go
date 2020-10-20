package medium

// Use Left and Right product lists:
// Time Complexity: O(3n) = O(n)
// Space Complexity:O(2n) = O(n)
// Runtime: 28 ms, faster than 29.00%
// Memory Usage: 6.6 MB, less than 37.13%
func productExceptSelf(nums []int) []int {
	size := len(nums)
	left := make([]int, size)
	right := make([]int, size)
	left[0] = 1
	for i := 1; i < size; i++ {
		left[i] = nums[i-1] * left[i-1]
	}
	right[size-1] = 1
	for i := size - 2; i >= 0; i-- {
		right[i] = nums[i+1] * right[i+1]
	}

	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = left[i] * right[i]
	}

	return result
}

// Optimize Left and Right product lists:
// Time Complexity: O(2n) = O(n)
// Space Complexity:O(1)
// Runtime: 28 ms, faster than 29.00%
// Memory Usage: 6.9 MB, less than 37.13%
func productExceptSelf2(nums []int) []int {
	size := len(nums)
	left := make([]int, size)

	left[0] = 1
	for i := 1; i < size; i++ {
		left[i] = nums[i-1] * left[i-1]
	}

	tmp := 1
	for i := size - 2; i >= 0; i-- {
		tmp *= nums[i+1]
		left[i] = tmp * left[i]
	}

	return left
}
