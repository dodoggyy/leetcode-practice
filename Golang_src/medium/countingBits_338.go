package medium

// Use Bit Manipulation:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 4 ms, faster than 92.41%
// Memory Usage: 4.4 MB, less than 33.79%
func countBits(num int) []int {
	result := make([]int, num+1)

	for i := 1; i <= num; i++ {
		result[i] = result[i>>1] + i&1
	}

	return result
}
