package easy

// Use hashset:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 5 ms, faster than 65.33%
// Memory Usage: 3.65 MB, less than 81.33%
func destCity(paths [][]string) string {
	hashset := map[string]bool{}
	for _, v := range paths {
		hashset[v[0]] = true
	}

	for _, v := range paths {
		if !hashset[v[1]] {
			return v[1]
		}
	}

	return ""
}
