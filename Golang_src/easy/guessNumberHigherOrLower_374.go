package easy

// Use binary search:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.9 MB, less than 88.15%
func guessNumber(n int) int {

	// mock function
	guess := func(num int) int {
		return num
	}

	l, r := 0, n

	for l <= r {
		mid := l + (r-l)/2
		switch guess(mid) {
		case -1:
			r = mid - 1
		case 1:
			l = mid + 1
		case 0:
			return mid
		}
	}

	return -1
}
