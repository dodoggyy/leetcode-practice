package medium

// Use simulation:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 48 ms, faster than 46.22%
// Memory Usage: 6.8 MB, less than 86.64%
func predictPartyVictory(senate string) string {
	r, d := []int{}, []int{}
	for i := range senate {
		if senate[i] == 'R' {
			r = append(r, i)
		} else {
			d = append(d, i)
		}
	}
	for len(r) > 0 && len(d) > 0 {
		if r[0] < d[0] {
			r = append(r, r[0]+len(senate))
		} else {
			d = append(d, d[0]+len(senate))
		}
		r = r[1:]
		d = d[1:]
	}
	if len(r) > 0 {
		return "Radiant"
	}
	return "Dire"
}
