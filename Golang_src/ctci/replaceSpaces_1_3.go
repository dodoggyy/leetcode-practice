package ctci

import "fmt"

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
func replaceSpaces(str []byte, trueLength int) string {
	spaceCount := 0

	for i := 0; i < trueLength; i++ {
		if str[i] == ' ' {
			spaceCount++
		}
	}

	index := trueLength + 2*spaceCount

	for i := trueLength - 1; i >= 0; i-- {
		if str[i] == ' ' {
			str[index-1] = '0'
			str[index-2] = '2'
			str[index-3] = '%'
			index -= 3
		} else {
			str[index-1] = str[i]
			index--
		}
	}

	return string(str)
}

func test() {
	fmt.Println(replaceSpaces([]byte("Mr John Smith    "), 13))
}
