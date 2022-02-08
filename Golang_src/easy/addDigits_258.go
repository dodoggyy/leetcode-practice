package easy

// Use recursive:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 7 ms, faster than 9.72%
// Memory Usage: 2.2 MB, less than 75.00%
func addDigits(num int) int {
	for num/10 != 0 {
		tmp := num
		num = 0
		for tmp != 0 {
			num += tmp % 10
			tmp /= 10
		}
	}

	return num
}

// Use math of digital root:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 90.28%
func addDigits2(num int) int {
	return (num-1)%9 + 1
}
