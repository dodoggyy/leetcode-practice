package medium

// Use XOR:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 107 ms, faster than 67.31%
// Memory Usage: 8.48 MB, less than 88.46%
func findArray(pref []int) []int {
	result := make([]int, len(pref))
	result[0] = pref[0]
	for i := range pref {
		if i > 0 {
			result[i] = pref[i] ^ pref[i-1]
		}
	}
	return result
}
