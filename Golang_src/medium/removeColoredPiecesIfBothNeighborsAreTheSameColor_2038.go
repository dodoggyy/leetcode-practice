package medium

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 13 ms, faster than 54.55%
// Memory Usage: 6.38 MB, less than 100.00%
func winnerOfGame(colors string) bool {
	cnt := [2]int{}
	cur := 'C'
	tmp := 0
	for _, ch := range colors {
		if ch == cur {
			tmp++
			if tmp >= 3 {
				cnt[ch-'A']++
			}
		} else {
			tmp = 1
			cur = ch
		}
	}

	return cnt[0] > cnt[1]
}
