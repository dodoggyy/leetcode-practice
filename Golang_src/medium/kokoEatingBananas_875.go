package medium

// Use binary search:
// m: max
// n: piles size
// Time Complexity: O(nlogm)
// Space Complexity:O(1)
// Runtime: 19 ms, faster than 95.87%
// Memory Usage: 7.09 MB, less than 26.38%
func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, v := range piles {
		if v > max {
			max = v
		}
	}

	helper := func(x int) bool {
		time := 0
		for _, v := range piles {
			time += ((v - 1) / x) + 1
		}

		return time <= h
	}

	l, r := 1, max

	for l < r {
		mid := l + (r-l)/2
		if helper(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
