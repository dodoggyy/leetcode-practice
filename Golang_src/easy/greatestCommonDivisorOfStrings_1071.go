package easy

// Use math:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 2 ms, faster than 67.97%
// Memory Usage: 2.48 MB, less than 61.72%
func gcdOfStrings(str1 string, str2 string) string {
	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}

	if str1+str2 != str2+str1 {
		return ""
	}

	return str1[:gcd(len(str1), len(str2))]
}
