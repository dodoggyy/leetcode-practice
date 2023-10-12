package medium

// Use 2D array simulation:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 8 ms, faster than 65.62%
// Memory Usage: 6.57 MB, less than 62.92%
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	arr := make([][]byte, numRows)
	cur := 0
	flag := -1
	for i := range s {
		arr[cur] = append(arr[cur], s[i])
		if cur == 0 || cur == numRows-1 {
			flag = -flag
		}
		cur += flag
	}

	result := []byte{}

	for i := range arr {
		result = append(result, arr[i]...)
	}

	return string(result)
}
