package medium

// Use binary search:
// Time Complexity: O(n*log(r-l))
// Space Complexity:O(1)
// Runtime: 36 ms, faster than 51.90%
// Memory Usage: 7.5 MB, less than 10.13%
func kthSmallest378(matrix [][]int, k int) int {
	n := len(matrix)

	// check target n-th position
	// O(n)
	checkCnt := func(target int) int {
		// start from left bottom (n-1, 0)
		i, j := n-1, 0
		cnt := 0

		for i >= 0 && j < n {
			if matrix[i][j] <= target {
				cnt += i + 1
				j++
			} else {
				i--
			}
		}

		return cnt
	}

	left, right := matrix[0][0], matrix[n-1][n-1]

	// O(r - l)
	for left <= right {
		mid := left + (right-left)/2
		cnt := checkCnt(mid)
		if cnt < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
