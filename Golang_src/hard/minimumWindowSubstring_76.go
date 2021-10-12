package hard

import "math"

// Use sliding window:
// Time Complexity: O(C * |s| + |t|)
// Space Complexity:O(C)
// C: character set
// Runtime: 16 ms, faster than 89.02%
// Memory Usage: 6.7 MB, less than 99.19%
func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	need := map[byte]int{}

	// keep checked string count
	for i := range t {
		need[t[i]]++
	}

	l, r := 0, 0          // left index, right index
	size := math.MaxInt32 // current result
	cnt := len(t)         // current need characters
	start := 0            // result start point

	for r < len(s) {
		// check right index ch is in string t or not
		if need[s[r]] > 0 {
			cnt--
		}
		need[s[r]]--

		// cnt == 0 means substring satisfied cover all string t
		if cnt == 0 {
			// need[s[l]] < 0 means left index ch is not in string t,
			// so that can remove it
			for l < r && need[s[l]] < 0 {
				need[s[l]]++
				l++
			}

			// update result if substring size < `size`
			if r-l+1 < size {
				size = r - l + 1
				start = l
			}

			// left index shift & it may lack 1 char in need map
			// so that have to upate cnt + 1
			need[s[l]]++
			l++
			cnt++
		}

		r++
	}

	if size == math.MaxInt32 {
		return ""
	}

	return s[start : start+size]
}
