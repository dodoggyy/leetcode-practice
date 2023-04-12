package hard

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 4 ms, faster than 96.84%
// Memory Usage: 3.9 MB, less than 34.18%
func longestConsecutive(nums []int) int {
	hashmap := make(map[int]int)
	result := 0

	for _, v := range nums {
		hashmap[v]++
	}

	for k := range hashmap {
		if _, exist := hashmap[k-1]; !exist {
			curNum := k
			tmpResult := 1

			_, exist2 := hashmap[curNum+1]
			for exist2 {
				curNum++
				tmpResult++
				_, exist2 = hashmap[curNum+1]
			}

			result = max(result, tmpResult)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 71 ms, faster than 79.6%
// Memory Usage: 10.2 MB, less than 72.89%
func longestConsecutive2(nums []int) int {
	hashset := map[int]bool{}

	for _, v := range nums {
		hashset[v] = true
	}
	result := 0

	for k := range hashset {
		if !hashset[k-1] {
			cur := k
			tmp := 1
			for hashset[cur+1] {
				cur++
				tmp++
			}
			if result < tmp {
				result = tmp
			}
		}
	}

	return result
}
