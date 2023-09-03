package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 4 ms, faster than 67.01%
// Memory Usage: 3.06 MB, less than 82.47%
func finalString(s string) string {
	result := []byte{}

	reverseStr := func() {
		for i := 0; i < len(result)/2; i++ {
			result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
		}
	}

	for i := range s {
		switch s[i] {
		case 'i':
			reverseStr()
		default:
			result = append(result, s[i])
		}
	}

	return string(result)
}
