package easy

// Use Array count:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 58.49%
// Memory Usage: 3.4 MB, less than 14.15%
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	countS, countT := [256]int{}, [256]int{}

	for i := range s {
		if countS[s[i]] != countT[t[i]] {
			return false
		}
		countS[s[i]] = (i + 1)
		countT[t[i]] = (i + 1)

	}

	return true
}
