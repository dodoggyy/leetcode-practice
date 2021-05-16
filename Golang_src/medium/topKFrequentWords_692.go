package medium

import "sort"

// Use bucket sort:
// Time Complexity: O(nlogn)
// Space Complexity:O(n)
// Runtime: 8 ms, faster than 27.94%
// Memory Usage: 4.5 MB, less than 57.35%
func topKFrequentWord(words []string, k int) []string {
	hashmap := map[string]int{}
	result := []string{}

	for _, word := range words {
		hashmap[word]++
	}

	bucket := map[int][]string{}

	for record, cnt := range hashmap {
		if value, ok := bucket[cnt]; ok {
			value = append(value, record)
			bucket[cnt] = value
		} else {
			bucket[cnt] = []string{record}
		}
	}

	curCnt := len(words)
	for k > len(result) {
		if val, ok := bucket[curCnt]; ok {
			sort.Strings(val)

			result = append(result, val...)
		}
		curCnt--
	}
	if len(result) > k {
		result = result[:k]
	}

	return result
}
