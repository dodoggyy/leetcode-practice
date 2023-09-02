package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.27 MB, less than 100.00%
func canBeEqual(s1 string, s2 string) bool {
	tmp := []byte(s1)

	for i := range tmp {
		if tmp[i] != s2[i] {
			j := (i + 2) % len(tmp)
			tmp[i], tmp[j] = tmp[j], tmp[i]
		}
	}

	return string(tmp) == s2
}
