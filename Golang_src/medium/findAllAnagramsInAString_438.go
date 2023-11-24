package medium

// Use sliding window:
// n: length of s
// m: length of p
// Σ: 26 alphabet
// Time Complexity: O(n+m+Σ)
// Space Complexity:O(Σ)
// Runtime: 6 ms, faster than 81.94%
// Memory Usage: 4.99 MB, less than 96.66%
func findAnagrams(s string, p string) []int {
	result := []int{}
	if len(s) < len(p) {
		return result
	}
	cntS, cntP := [26]int{}, [26]int{}
	for i := range p {
		cntP[p[i]-'a']++
	}
	for i := range s {
		cntS[s[i]-'a']++

		if i >= len(p) {
			cntS[s[i-len(p)]-'a']--
		}

		if cntP == cntS {
			result = append(result, i-len(p)+1)
		}
	}

	return result
}
