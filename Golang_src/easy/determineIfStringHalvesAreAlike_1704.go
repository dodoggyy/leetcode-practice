package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 2 ms, faster than 72.09%
// Memory Usage: 2.22 MB, less than 20.93%
func halvesAreAlike(s string) bool {
	isVowel := func(ch byte) bool {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		case 'A', 'E', 'I', 'O', 'U':
			return true
		}

		return false
	}

	cnt1, cnt2 := 0, 0
	for i := 0; i < len(s)/2; i++ {
		if isVowel(s[i]) {
			cnt1++
		}
		if isVowel(s[len(s)/2+i]) {
			cnt2++
		}
	}

	return cnt1 == cnt2
}
