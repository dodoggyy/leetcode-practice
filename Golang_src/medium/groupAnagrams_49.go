package medium

import "sort"

// Use Categorize by Sorted String:
// n for length of strs
// k for max length string in strs
// Time Complexity: O(n*klogk)
// Space Complexity:O(n*k)
// Runtime: 24 ms, faster than 85.57%
// Memory Usage: 7.9 MB, less than 36.87%
func groupAnagrams(strs []string) [][]string {
	hashmap := map[string][]string{}
	for _, str := range strs {
		tmp := []byte(str)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})

		sortStr := string(tmp)
		hashmap[sortStr] = append(hashmap[sortStr], str)
	}
	result := make([][]string, 0, len(hashmap))
	for _, v := range hashmap {
		result = append(result, v)
	}

	return result
}

// Use Categorize by Count:
// n for length of strs
// k for max length string in strs
// Time Complexity: O(n*k)
// Space Complexity:O(n*k)
// Runtime: 24 ms, faster than 85.57%
// Memory Usage: 7.4 MB, less than 64.33%
func groupAnagrams2(strs []string) [][]string {
	hashmap := map[[26]int][]string{}

	for _, str := range strs {
		count := [26]int{}
		for _, ch := range str {
			count[ch-'a']++
		}
		hashmap[count] = append(hashmap[count], str)
	}
	result := make([][]string, 0, len(hashmap))
	for _, v := range hashmap {
		result = append(result, v)
	}

	return result
}
