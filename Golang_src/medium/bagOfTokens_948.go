package medium

import "sort"

// Use sort + two pointers:
// Time Complexity: O(nlogn)
// Space Complexity:O(logn)
// Runtime: 9 ms, faster than 14.29%
// Memory Usage: 3 MB, less than 42.86%
func bagOfTokensScore(tokens []int, power int) int {
	result := 0
	sort.Ints(tokens)

	point := 0
	l, r := 0, len(tokens)-1
	for l <= r && (tokens[l] <= power || point > 0) {
		for l <= r && tokens[l] <= power {
			point++
			power -= tokens[l]
			l++
		}

		result = max(result, point)

		if l <= r && point > 0 {
			point--
			power += tokens[r]
			r--
		}
	}

	return result
}
