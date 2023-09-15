package medium

import "strconv"

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 7 ms, faster than 12.19%
// Memory Usage: 6.61 MB, less than 5.19%
func compress(chars []byte) int {
	result := ""
	cnt := 1

	cur := chars[0]
	for i := 1; i < len(chars); i++ {
		if chars[i] != cur {
			result += string(cur)
			if cnt > 1 {
				result += strconv.Itoa(cnt)
			}
			cnt = 1
			cur = chars[i]

		} else {
			cnt++
		}
	}

	result += string(cur)
	if cnt > 1 {
		result += strconv.Itoa(cnt)
	}

	for i := range result {
		chars[i] = result[i]
	}

	return len(result)
}

// Use optimize iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 5 ms, faster than 72.46%
// Memory Usage: 2.84 MB, less than 25.96%
func compress2(chars []byte) int {
	cnt := 1
	idx := 0

	cur := chars[0]
	for i := 1; i < len(chars); i++ {
		if chars[i] != cur {
			chars[idx] = cur
			idx++
			if cnt > 1 {
				digits := strconv.Itoa(cnt)
				for i := range digits {
					chars[idx] = digits[i]
					idx++
				}
			}
			cnt = 1
			cur = chars[i]

		} else {
			cnt++
		}
	}

	chars[idx] = cur
	idx++
	if cnt > 1 {
		digits := strconv.Itoa(cnt)
		for i := range digits {
			chars[idx] = digits[i]
			idx++
		}
	}

	return idx
}
