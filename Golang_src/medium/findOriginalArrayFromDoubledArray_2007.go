package medium

import "sort"

// Use counting sort:
// Time Complexity: O(n+k)
// Space Complexity:O(k)
// Runtime: 176 ms, faster than 97.37%
// Memory Usage: 10.6 MB, less than 81.58%
func findOriginalArray2(changed []int) []int {
	result := []int{}
	if len(changed)%2 != 0 {
		return result
	}

	maxNum := 0
	for _, v := range changed {
		if v > maxNum {
			maxNum = v
		}
	}

	freq := make([]int, maxNum*2+1)
	for _, v := range changed {
		freq[v]++
	}

	for i := 0; i <= maxNum; i++ {
		if freq[i] > 0 {
			freq[i]--

			twice := i * 2
			if freq[twice] > 0 {
				freq[twice]--
				result = append(result, i)
				i--
			} else {
				return []int{}
			}
		}
	}

	return result
}

// Use sort + hashmap:
// Time Complexity: O(nolgn)
// Space Complexity:O(n)
// Runtime: 633 ms, faster than 31.58%
// Memory Usage: 11.8 MB, less than 55.26%
func findOriginalArray(changed []int) []int {
	result := []int{}
	if len(changed)%2 != 0 {
		return result
	}

	sort.Ints(changed)

	hashmap := map[int]int{}

	for _, v := range changed {
		if hashmap[v] == 0 {
			hashmap[v*2]++
			result = append(result, v)
		} else {
			hashmap[v]--
			if hashmap[v] == 0 {
				delete(hashmap, v)
			}
		}
	}

	if len(hashmap) != 0 {
		return []int{}
	}

	return result
}
