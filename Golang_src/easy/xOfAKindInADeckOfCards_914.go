package easy

// Use highest common factor:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100%
// Memory Usage: 2.4 MB, less than 23.81%
func hasGroupsSizeX(deck []int) bool {
	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if a == 0 {
			return b
		}
		return gcd(b%a, a)
	}

	hashmap := map[int]int{}
	for _, v := range deck {
		hashmap[v]++
	}

	tmp := 0
	hadSet := false
	for _, v := range hashmap {
		if !hadSet {
			tmp = v
			hadSet = true
		}
		if hadSet {
			tmp = gcd(tmp, v)
		}
	}

	return tmp >= 2
}
