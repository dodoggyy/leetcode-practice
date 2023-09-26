package medium

// Use greedy + monotonic stack:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.1 MB, less than 43.33%
func removeDuplicateLetters(s string) string {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	stack := []byte{}
	remainInAns := [26]bool{}
	for i := range s {
		ch := s[i]

		if !remainInAns[ch-'a'] {
			for len(stack) > 0 && stack[len(stack)-1] > s[i] {
				if cnt[stack[len(stack)-1]-'a'] == 0 {
					break
				}
				remainInAns[stack[len(stack)-1]-'a'] = false
				stack = stack[:len(stack)-1]
			}

			stack = append(stack, ch)
			remainInAns[ch-'a'] = true

		}

		cnt[ch-'a']--
	}

	return string(stack)
}
