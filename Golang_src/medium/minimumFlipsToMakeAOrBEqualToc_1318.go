package medium

// Use bit mainupation:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 1 ms, faster than 81.82%
// Memory Usage: 1.9 MB, less than 100.00%
func minFlips(a int, b int, c int) int {
	result := 0

	for i := 0; i <= 31; i++ {
		bitA := (a >> i) & 1
		bitB := (b >> i) & 1
		bitC := (c >> i) & 1

		if bitC == 0 {
			if bitA == 1 {
				result++
			}
			if bitB == 1 {
				result++
			}
		} else {
			if bitA == 0 && bitB == 0 {
				result++
			}
		}
	}

	return result
}
