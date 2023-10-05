package medium

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 8 ms, faster than 92.20%
// Memory Usage: 5.04 MB, less than 43.97%
func majorityElement(nums []int) []int {
	result := []int{}
	target := len(nums) / 3
	hashmap := map[int]int{}
	for _, v := range nums {
		hashmap[v]++
	}

	for k, v := range hashmap {
		if v > target {
			result = append(result, k)
		}
	}

	return result
}

// Use Boyer-Moore Voting Algorithm:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 12 ms, faster than 57.45%
// Memory Usage: 5.04 MB, less than 43.97%
func majorityElement2(nums []int) []int {
	result := []int{}
	element1, element2 := 0, 0
	cnt1, cnt2 := 0, 0
	target := len(nums) / 3

	for _, v := range nums {
		if cnt1 > 0 && element1 == v {
			cnt1++
		} else if cnt2 > 0 && element2 == v {
			cnt2++
		} else if cnt1 == 0 {
			element1 = v
			cnt1++
		} else if cnt2 == 0 {
			element2 = v
			cnt2++
		} else {
			cnt1--
			cnt2--
		}
	}

	cnt1, cnt2 = 0, 0
	for _, v := range nums {
		if v == element1 {
			cnt1++
		} else if v == element2 {
			cnt2++
		}
	}

	if cnt1 > target {
		result = append(result, element1)
	}
	if cnt2 > target {
		result = append(result, element2)
	}

	return result
}
