package easy

import "strconv"

// Use iterative:
// Time Complexity: O(max(m,n))
// Space Complexity:O(max(m,n))
// m : length of num1
// n : length of num2
// Runtime: 8 ms, faster than 44.51%
// Memory Usage: 7.1 MB, less than 7.14%
func addStrings(num1 string, num2 string) string {

	result := ""

	str1 := []rune(num1)
	str2 := []rune(num2)

	index1 := len(str1) - 1
	index2 := len(str2) - 1

	carry := 0

	for index1 >= 0 || index2 >= 0 || carry != 0 {
		if index1 >= 0 {
			carry += int(str1[index1] - '0')
			index1--
		}
		if index2 >= 0 {
			carry += int(str2[index2] - '0')
			index2--
		}
		result = strconv.Itoa(carry%10) + result
		carry /= 10

	}

	return result
}
