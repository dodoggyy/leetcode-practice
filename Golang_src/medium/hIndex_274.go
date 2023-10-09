package medium

// Use counting:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 2 ms, faster than 77.88%
// Memory Usage: 2.33 MB, less than 11.91%
func hIndex(citations []int) int {
	size := len(citations)
	cnts := make([]int, size+1)
	for _, v := range citations {
		if v >= size {
			cnts[size]++
		} else {
			cnts[v]++
		}
	}

	total := 0
	for i := size; i >= 0; i-- {
		total += cnts[i]
		if total >= i {
			return i
		}
	}

	return 0
}

// Use binary search:
// Time Complexity: O(nlogn)
// Space Complexity:O(1)
// Runtime: 2 ms, faster than 77.88%
// Memory Usage: 2.24 MB, less than 64.84%
func hIndex2(citations []int) int {
	l, r := 0, len(citations)
	for l < r {
		// +1 prevent infinite loop
		mid := (l + r + 1) / 2
		cnt := 0
		for _, v := range citations {
			if v >= mid {
				cnt++
			}
		}

		if cnt >= mid {
			l = mid
		} else {
			r = mid - 1
		}
	}

	return l
}
