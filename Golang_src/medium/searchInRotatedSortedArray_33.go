package medium

// Use binary serach:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.6 MB, less than 100.00%
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] { // [l,mid] is ascending order
			// left < target < mid
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // [mid, right] is ascending order
			// mid < target < right
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
