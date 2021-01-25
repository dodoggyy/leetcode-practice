package easy

// Use reverse number:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 93.17%
// Memory Usage: 5.2 MB, less than 44.36%
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	tmp, copy := 0, x

	for copy != 0 {
		tmp = tmp*10 + copy%10
		copy /= 10
	}

	return tmp == x
}
