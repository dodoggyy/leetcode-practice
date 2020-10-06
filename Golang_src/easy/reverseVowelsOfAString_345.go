package easy

// Use two pointers:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 84.17%
// Memory Usage: 4.5 MB, less than 58.27%
func reverseVowels(s string) string {
	str := []rune(s)
	left, right := 0, len(str)-1

	for left < right {
		if !isVowel(str[left]) {
			left++
			continue
		}
		if !isVowel(str[right]) {
			right--
			continue
		}

		str[left], str[right] = str[right], str[left]

		left++
		right--
	}

	return string(str)
}

func isVowel(ch rune) (isVowel bool) {
	if ch < 'a' {
		ch += 32
	}
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}
