package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 6 ms, faster than 16.94%
// Memory Usage: 2.13 MB, less than 25.81%
func furthestDistanceFromOrigin(moves string) int {
	result := 0
	tmp := 0

	for _, ch := range moves {
		switch ch {
		case 'L':
			result++
		case 'R':
			result--
		case '_':
			tmp++
		}
	}

	abs := func(a int) int {
		if a > 0 {
			return a
		}
		return -a
	}

	return abs(result) + tmp
}
