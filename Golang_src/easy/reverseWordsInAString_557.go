package easy

// Use extra space:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 56 ms, faster than 11.53%
// Memory Usage: 8.3 MB, less than 13.47%
func reverseWords(s string) string {
	reverse := func(str []byte) string {
		size := len(str)
		for i := 0; i < size/2; i++ {
			str[i], str[size-1-i] = str[size-1-i], str[i]
		}

		return string(str)
	}

	result := ""

	tmp := ""
	for i := range s {
		if s[i] == ' ' {
			if len(tmp) == 0 {
				result += " "
				continue
			}
			result += reverse([]byte(tmp))
			result += " "
			tmp = ""
		} else {
			tmp += string(s[i])
		}
	}
	if len(tmp) > 0 {
		result += reverse([]byte(tmp))
	}

	return result
}

// Use in space:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 62.82%
//Memory Usage: 6.4 MB, less than 92.69%
func reverseWords2(s string) string {
	str := []rune(s)
	i := 0

	for i < len(str) {
		start := i
		for i < len(str) && str[i] != ' ' {
			i++
		}

		left, right := start, i-1
		for left < right {
			str[left], str[right] = str[right], str[left]
			left++
			right--
		}

		i++
	}

	return string(str)
}
