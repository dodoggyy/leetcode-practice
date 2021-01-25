package easy

// Use DP:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 33.77%
func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	result := 0
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1

	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			arr[i] = arr[i/2]
		} else {
			arr[i] = arr[i/2] + arr[i/2+1]
		}
		result = max(result, arr[i])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
