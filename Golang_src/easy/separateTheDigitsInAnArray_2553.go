package easy

import "strconv"

// Use simulation:
// Time Complexity: O(d*n)
// Space Complexity:O(d*n)
// Runtime: 9 ms, faster than 30.39%
// Memory Usage: 5.7 MB, less than 60.77%
func separateDigits(nums []int) []int {
	result := []int{}

	for _, v := range nums {
		str := strconv.Itoa(v)
		for _, v := range str {
			digit, _ := strconv.Atoi(string(v))
			result = append(result, digit)
		}
	}

	return result
}
