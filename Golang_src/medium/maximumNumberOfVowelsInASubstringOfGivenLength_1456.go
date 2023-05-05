package medium

// Use sliding window:
// Time Complexity: O(|S|)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 97.44%
// Memory Usage: 4.7 MB, less than 96.15%
func maxVowels(s string, k int) int {
	l, r := 0, 0
	size := len(s)
	result := 0

	isVowel := func(ch byte) bool {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		}

		return false
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	tmp := 0
	for r < size {
		if result == k {
			return result
		}
		if isVowel(s[r]) {
			tmp++
		}
		if r-l+1 > k {
			if isVowel(s[l]) {
				tmp--
			}
			l++
		}
		r++

		result = max(result, tmp)
	}

	return result
}
