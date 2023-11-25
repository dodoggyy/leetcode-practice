package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 3.42 MB, less than 70.09%
func findMinimumOperations(s1 string, s2 string, s3 string) int {
	idx1, idx2, idx3 := 0, 0, 0
	size1, size2, size3 := len(s1), len(s2), len(s3)

	for idx1 < size1 && idx2 < size2 && idx3 < size3 && s1[idx1] == s2[idx2] && s2[idx2] == s3[idx3] {
		idx1++
		idx2++
		idx3++
	}

	if idx1 == 0 {
		return -1
	}

	return size1 + size2 + size3 - idx1 - idx2 - idx3
}
