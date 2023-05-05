package easy

// Use iterative times:
// Time Complexity: O(n^(1/2))
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.9 MB, less than 100.00%
func isPerfectSquare(num int) bool {
	for i := 1; i <= num/i; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}

// Use arithmetic sequence:
// 1, 4, 9, 16, ..., n^2
// arithmetic sequence: 3, 5 , 7, ..., (2n + 1)
// Time Complexity: O(n^(1/2))
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.9 MB, less than 100.00%
func isPerfectSquare2(num int) bool {
	subNum := 1
	for i := 1; i <= num; i += subNum {
		if i == num {
			return true
		}
		subNum += 2
	}
	return false
}

// Use binary search:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 2 ms, faster than 47.49%
// Memory Usage: 1.9 MB, less than 97.77%
func isPerfectSquare3(num int) bool {
	l, r := 1, num
	for l <= r {
		mid := l + (r-l)/2
		cur := mid * mid
		if cur == num {
			return true
		} else if cur < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
