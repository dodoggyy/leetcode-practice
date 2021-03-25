package easy

// Use binary search:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// untime: 4 ms, faster than 85.89%
// Memory Usage: 3.1 MB, less than 100%
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
