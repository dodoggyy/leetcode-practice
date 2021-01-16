package medium

// Use binary search:
// Time Complexity: O(logn)~O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 91.86%
// Memory Usage: 3.4 MB, less than 7.12%
func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return true
		}
		// diff from LC No.33
		if nums[left] == nums[mid] {
			left++
			continue
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

	return false
}
