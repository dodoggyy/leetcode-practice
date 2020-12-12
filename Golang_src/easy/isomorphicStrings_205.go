package easy

// Use Array count:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 58.49%
// Memory Usage: 3.4 MB, less than 14.15%
func isIsomorphic(s string, t string) bool {
	strS := []rune(s)
	strT := []rune(t)

	countS, countT := [256]int{}, [256]int{}

	length := len(strS)

	for i := 0; i < length; i++ {
		if countS[strS[i]] != countT[strT[i]] {
			return false
		}
		countS[strS[i]] = (i + 1)
		countT[strT[i]] = (i + 1)

	}

	return true
}
