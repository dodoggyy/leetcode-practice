package ctci

import "strconv"

// Use iterative
// Time Complexity: O(n)
// Space Complexity:O(1)
func BasicCompress(input string) string {
	if len(input) < 2 {
		return input
	}

	cnt := 1
	result := ""
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			cnt++
		} else {
			result = result + string(input[i-1]) + strconv.Itoa(cnt)
			cnt = 1
		}
	}

	result = result + string(input[len(input)-1]) + strconv.Itoa(cnt)

	if len(result) > len(input) {
		return input
	} else {
		return result
	}
}
