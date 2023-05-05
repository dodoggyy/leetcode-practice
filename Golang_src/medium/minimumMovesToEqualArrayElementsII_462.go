package medium

import (
	"sort"
)

// Use sorting median:
// Time Complexity: O(nlogn)
// Space Complexity:O(1)
// Runtime: 12 ms, faster than 76.32%
// Memory Usage: 4.5 MB, less than 10.53%
func minMoves2(nums []int) int {
	result := 0

	sort.Ints(nums)
	median := nums[len(nums)/2]

	for _, v := range nums {
		result += abs(median - v)
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Use quick select:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: Time Limit Exceeded
// Memory Usage: Time Limit Exceeded
func minMoves22(nums []int) int {
	result := 0
	middle := len(nums) / 2
	medium := findMedium(nums, middle)

	for _, v := range nums {
		result += abs(medium - v)
	}

	return result
}

func findMedium(nums []int, index int) int {
	start := 0
	end := len(nums) - 1

	for {
		if start == end {
			return nums[start]
		}

		pivot := nums[start]
		l := start + 1
		r := end

		for {
			for r > start && nums[r] <= pivot {
				r--
			}
			for l < end && nums[l] >= pivot {
				l++
			}

			if l < r {
				nums[l], nums[r] = nums[r], nums[l]
			} else {
				nums[start], nums[r] = nums[r], nums[start]
			}
			break
		}

		if r == index {
			return nums[r]
		} else if r > index {
			end = l - 1
		} else {
			start = r + 1
		}
	}
}

// Use quick select 2:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 149 ms, faster than 8.70%
// Memory Usage: 5.2 MB, less than 100%
func minMoves3(nums []int) int {
	size := len(nums)

	var qs func(l, r, k int) int
	qs = func(l, r, k int) int {
		i := l - 1
		for j := l; j < r; j++ {
			if nums[j] < nums[r] {
				i++
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
		i++
		nums[i], nums[r] = nums[r], nums[i]

		if i == k {
			return nums[i]
		} else if i < k {
			return qs(i+1, r, k)
		} else {
			return qs(l, i-1, k)
		}
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	mid := qs(0, size-1, size/2)

	result := 0
	for _, v := range nums {
		result += abs(mid - v)
	}

	return result
}
