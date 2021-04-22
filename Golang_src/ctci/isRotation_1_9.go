package ctci

import "strings"

// Use concat string compare:
// Time Complexity: O(n)
// Space Complexity:O(n)
func IsRotation(input1, input2 string) bool {
	if len(input1) != len(input2) {
		return false
	}

	newStr := input1 + input1

	return strings.Contains(newStr, input2)
}
