package medium

// Use bucket sort:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 12 ms, faster than 86.24%
// Memory Usage: 5.7 MB, less than 67.58%
func topKFrequent(nums []int, k int) []int {
	result := []int{}
	size := len(nums)
	if size < k {
		return result
	}

	hashmap := map[int]int{}

	for _, v := range nums {
		if val, exist := hashmap[v]; exist {
			hashmap[v] = val + 1
		} else {
			hashmap[v] = 1
		}
	}

	bucket := map[int][]int{}

	for key, val := range hashmap {
		if v, exist := bucket[val]; exist {
			v = append(v, key)
			bucket[val] = v
		} else {
			bucket[val] = []int{key}
		}
	}

	idx := size
	for len(result) < k {
		if v, exist := bucket[idx]; exist {
			result = append(result, v...)
		}
		idx--
	}

	return result
}
