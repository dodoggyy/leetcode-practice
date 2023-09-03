package easy

import (
	"strconv"
)

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 76 ms, faster than 100.00%
// Memory Usage: 8.2 MB, less than 100.00%
func countSymmetricIntegers(low int, high int) int {
	result := 0

	isSym := func(n int) bool {
		str := strconv.Itoa(n)

		if len(str)%2 != 0 {
			return false
		}

		tmp := 0

		for i := 0; i < len(str)/2; i++ {
			left, _ := strconv.Atoi(string(str[i]))
			right, _ := strconv.Atoi(string(str[len(str)-1-i]))

			tmp += left
			tmp -= right
		}

		if tmp != 0 {
			return false
		}

		return true
	}

	for i := low; i <= high; i++ {
		if isSym(i) {
			result++
		}
	}

	return result
}
