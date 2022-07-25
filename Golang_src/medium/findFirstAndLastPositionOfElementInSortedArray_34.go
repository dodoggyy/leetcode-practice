package medium

// Use binary search:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 7 ms, faster than 82.56%
// Memory Usage: 3.9 MB, less than 75.83%
func searchRange(nums []int, target int) []int {
	var binarySearch func(l, r int, findLow bool) int
	binarySearch = func(l, r int, findLow bool) int {
		result := -1

		for l <= r {
			mid := l + (r-l)/2

			if findLow {
				if nums[mid] <= target {
					l = mid + 1
				} else {
					r = mid - 1
				}
			} else {
				if nums[mid] < target {
					l = mid + 1
				} else {
					r = mid - 1
				}
			}

			if nums[mid] == target {
				result = mid
			}
		}

		return result
	}

	return []int{binarySearch(0, len(nums)-1, false), binarySearch(0, len(nums)-1, true)}
}
