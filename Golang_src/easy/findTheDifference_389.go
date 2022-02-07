package easy

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(|Σ|)
// for |Σ| = 26
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.1 MB, less than 95.41%
func findTheDifference(s string, t string) byte {
	cnt := make([]int, 26)

	for _, v := range s {
		cnt[v-'a']++
	}

	for _, v := range t {
		cnt[v-'a']--
		if cnt[v-'a'] < 0 {
			return byte(v)
		}
	}

	return ' '
}

// Use XOR:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 3 ms, faster than 42.20%
// Memory Usage: 2.1 MB, less than 95.41%
func findTheDifference2(s string, t string) byte {
	var diff byte

	for i := range s {
		diff ^= s[i] ^ t[i]
	}

	diff ^= t[len(t)-1]

	return diff
}
