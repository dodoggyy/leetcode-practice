package medium

import "sort"

// Use sorting:
// Time Complexity: O(nlogn)
// Space Complexity:O(n)
// Runtime: 19 ms, faster than 89.19%
// Memory Usage: 6.49 MB, less than 100.00%
func sortVowels(s string) string {
	tmp := []byte{}
	str := []byte(s)

	isVowel := func(ch byte) bool {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		case 'A', 'E', 'I', 'O', 'U':
			return true
		}
		return false
	}

	for i := range s {
		if isVowel(s[i]) {
			tmp = append(tmp, s[i])
		}
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	idx := 0
	for i := range str {
		if isVowel(str[i]) {
			str[i] = tmp[idx]
			idx++
		}
	}

	return string(str)
}

// Use count:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 17 ms, faster than 89.19%
// Memory Usage: 7.81 MB, less than 40.54%
func sortVowels2(s string) string {
	str := []byte(s)
	idx := []int{}
	vowel := make([]int, 256)

	isVowel := func(ch byte) bool {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		case 'A', 'E', 'I', 'O', 'U':
			return true
		}
		return false
	}

	for i := range str {
		if isVowel(str[i]) {
			idx = append(idx, i)
			vowel[str[i]]++
		}
	}

	cur := 0
	for i, cnt := range vowel {
		if cnt != 0 {
			for j := 0; j < cnt; j++ {
				str[idx[cur]] = byte(i)
				cur++
			}
		}
	}

	return string(str)
}
