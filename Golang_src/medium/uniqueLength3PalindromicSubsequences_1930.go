package medium

// Use Pre-Compute First and Last Indices:
// Time Complexity: O(n+|Σ|)
// Space Complexity:O(|Σ|)
// Runtime: 42 ms, faster than 88.89%
// Memory Usage: 6.08 MB, less than 100.00%
func countPalindromicSubsequence(s string) int {
	result := 0

	pre, suf, has := [26]int{}, [26]int{}, [26]int{}
	for _, ch := range s[1:] {
		suf[ch-'a']++
	}
	for i := 1; i < len(s)-1; i++ {
		pre[s[i-1]-'a']++
		suf[s[i]-'a']--
		for j := 0; j < 26; j++ {
			if pre[j] > 0 && suf[j] > 0 {
				has[s[i]-'a'] |= 1 << j
			}
		}

	}

	cntBit := func(n int) int {
		cnt := 0
		for n > 0 {
			cnt++
			n &= n - 1
		}
		return cnt
	}

	for _, mask := range has {
		result += cntBit(mask)
	}

	return result
}
