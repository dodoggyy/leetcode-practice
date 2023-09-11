package medium

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 8 ms, faster than 57.14%
// Memory Usage: 5.38 MB, less than 53.57%
func groupThePeople(groupSizes []int) [][]int {
	result := [][]int{}

	hashmap := map[int][]int{}

	for i := range groupSizes {
		if v, ok := hashmap[groupSizes[i]]; ok {
			v = append(v, i)
			hashmap[groupSizes[i]] = v
		} else {
			hashmap[groupSizes[i]] = []int{i}
		}
	}

	for k, v := range hashmap {
		for i := 0; i < len(v); i += k {
			result = append(result, v[i:i+k])
		}
	}

	return result
}
