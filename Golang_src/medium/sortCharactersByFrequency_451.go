package medium

import "sort"

// Use hashmap and sort:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 57.81%
// Memory Usage: 6 MB, less than 26.56%
func frequencySort(s string) string {
	hashmap := map[rune]int{}
	chs := make([]rune, 0)

	for _, ch := range s {
		if value, exist := hashmap[ch]; exist {
			hashmap[ch] = value + 1
		} else {
			hashmap[ch] = 1
			chs = append(chs, ch)
		}
	}

	sort.Slice(chs, func(i, j int) bool {
		return hashmap[chs[i]] > hashmap[chs[j]]
	})

	result := []rune{}

	for _, v := range chs {
		freq := hashmap[v]
		for j := 0; j < freq; j++ {
			result = append(result, v)
		}
	}

	return string(result)
}
