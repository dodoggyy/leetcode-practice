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
