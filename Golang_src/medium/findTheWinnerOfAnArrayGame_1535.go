package medium

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 85 ms, faster than 92.86%
// Memory Usage: 7.94 MB, less than 100.00%
func getWinner(arr []int, k int) int {
	idx, cnt := 0, 0

	for cnt < k && idx < len(arr)-1 {
		if arr[idx] > arr[idx+1] {
			arr[idx+1] = arr[idx]
			cnt++
		} else {
			cnt = 1
		}
		idx++
	}
	return arr[idx]
}
