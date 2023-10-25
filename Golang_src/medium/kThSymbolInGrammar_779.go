package medium

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.91 MB, less than 92.50%
func kthGrammar(n int, k int) int {
	// 0
	// 01
	// 0110
	// 01101001
	// 0110100110010110

	if n == 1 {
		return 0
	}

	if k > 1<<(n-2) {
		return kthGrammar(n-1, k-1<<(n-2)) ^ 1
	}
	return kthGrammar(n-1, k)
}
