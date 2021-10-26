package medium

import "sort"

// Use build in sort:
// Time Complexity: O(nlogn)
// Space Complexity:O(logn)
// Runtime: 120 ms, faster than 69.62%
// Memory Usage: 8.2 MB, less than 51.88%
func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return calDistance(points[i]) < calDistance(points[j])
	})

	return points[:k]
}

func calDistance(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

// Use  quick sort:
// Time Complexity: O(nlogn)
// Space Complexity:O(logn)
// Runtime: 108 ms, faster than 84.64%
// Memory Usage: 8.2 MB, less than 51.88%
func kClosest2(points [][]int, k int) [][]int {
	var partition func(l, r int) int
	partition = func(l, r int) int {
		i := l - 1
		pivot := calDistance(points[r])
		for j := l; j < r; j++ {
			cur := calDistance(points[j])
			if cur < pivot {
				i++
				points[j], points[i] = points[i], points[j]
			}
		}
		i++
		points[r], points[i] = points[i], points[r]

		return i
	}

	var qsort func(l, r int)
	qsort = func(l, r int) {
		if l < r {
			q := partition(l, r)
			qsort(l, q-1)
			qsort(q+1, r)
		}
	}

	qsort(0, len(points)-1)

	return points[:k]
}

// Use  quick select:
// Time Complexity: O(n)
// Space Complexity:O(logn)
// Runtime: 92 ms, faster than 96.93%
// Memory Usage: 8.3 MB, less than 51.88%
func kClosest3(points [][]int, k int) [][]int {
	var qselect func(l, r int)
	qselect = func(l, r int) {
		if l == r {
			return
		}

		i := l - 1
		pivot := calDistance(points[r])
		for j := l; j < r; j++ {
			cur := calDistance(points[j])
			if cur < pivot {
				i++
				points[i], points[j] = points[j], points[i]
			}
		}
		i++
		points[i], points[r] = points[r], points[i]

		if i+1 == k {
			return
		} else if i+1 < k {
			qselect(i+1, r)
		} else {
			qselect(l, i-1)
		}
	}

	qselect(0, len(points)-1)

	return points[:k]
}
