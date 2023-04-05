package medium

import "sort"

// Use hashmap and sort:
// Time Complexity: O(nlogn)
// Space Complexity:O(n)
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

// Use Multiset and Bucket Sort:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 3 ms, faster than 94.78%
// Memory Usage: 5.8 MB, less than 41.74%
func frequencySor2(s string) string {
	hashmap := map[byte]int{}

	for i := range s {
		hashmap[s[i]]++
	}

	hashmap2 := map[int][]byte{}

	for k, v := range hashmap {
		if val, ok := hashmap2[v]; ok {
			val = append(val, k)
			hashmap2[v] = val
		} else {
			hashmap2[v] = []byte{k}
		}
	}

	result := []byte{}
	cur := len(s)
	for len(result) < len(s) {
		if val, ok := hashmap2[cur]; ok {
			for _, ch := range val {
				for j := 0; j < cur; j++ {
					result = append(result, ch)
				}
			}
		}
		cur--
	}

	return string(result)
}
