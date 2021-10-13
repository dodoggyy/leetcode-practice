package medium

// Use quick sort:
// Time Complexity: O(nlogn)
// Space Complexity:O(1)
// Runtime: 3223 ms, faster than 5.13%
// Memory Usage: 12.7 MB, less than 20.51%
func sortArray(nums []int) []int {
	partition := func(l, r int) int {
		i := l - 1
		for j := l; j < r; j++ {
			if nums[j] < nums[r] {
				i++
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
		i++
		nums[r], nums[i] = nums[i], nums[r]

		return i
	}

	var qsort func(l, r int)
	qsort = func(l, r int) {
		if l < r {
			q := partition(l, r)
			qsort(l, q-1)
			qsort(q+1, r)
		}
	}

	qsort(0, len(nums)-1)

	return nums
}
