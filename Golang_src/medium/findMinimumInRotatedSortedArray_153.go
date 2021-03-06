package medium

// Use binary search:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.5 MB, less than 18.69%
func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			// mid can be min
			right = mid
		} else {
			// mid cannot be min
			left = mid + 1
		}
	}
	return nums[left]
}
