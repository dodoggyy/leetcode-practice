package medium

import "sort"

// Use DP:
// Time Complexity: O(nlogn + n*L^2) = O(n*(logn+L^2))
// Space Complexity:O(n)
// L: word max length
// Runtime: 35 ms, faster than 57.56%
// Memory Usage: 6.9 MB, less than 31.40%
func longestStrChain(words []string) int {
	result := 0
	hashmap := map[string]int{}

	sort.Slice(words, func(a, b int) bool {
		return len(words[a]) < len(words[b])
	})

	for _, v := range words {
		tmp := 0
		for i := range v {
			tmp = max(tmp, hashmap[v[:i]+v[i+1:]]+1)
		}
		hashmap[v] = tmp
	}

	for _, v := range hashmap {
		if v > result {
			result = v
		}
	}

	return result
}
