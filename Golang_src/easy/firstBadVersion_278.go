package easy

func isBadVersion(version int) bool

// Use binary search:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.9 MB, less than 47.15%
func firstBadVersion(n int) int {
	l, r := 1, n
	for l < r {
		mid := l + (r-l)/2
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}
