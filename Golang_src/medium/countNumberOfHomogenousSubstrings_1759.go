package medium

// Use math:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 10 ms, faster than 77.78%
// Memory Usage: 6.22 MB, less than 85.19%
func countHomogenous(s string) int {
	result := 0
	mod := int(1e9 + 7)

	cur := s[0]
	tmp := 1
	for i := 1; i < len(s); i++ {
		if s[i] == cur {
			tmp++
		} else {
			result += (tmp + 1) * tmp / 2
			tmp = 1
			cur = s[i]
		}
	}
	result += (tmp + 1) * tmp / 2

	return result % mod
}

// a  1
// aa 1+2
// aaa 1 + 2 + 3
