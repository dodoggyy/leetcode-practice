package easy

import "sort"

// Use sorting:
// Time Complexity: O(mlogm)
// Space Complexity:O(m)
// Runtime: 11 ms, faster than 61.84%
// Memory Usage: 4.9 MB, less than 72.37%
func kWeakestRows(mat [][]int, k int) []int {
	m := len(mat)

	arr := make([]int, m)
	idx := make([]int, m)

	for i := range mat {
		cnt := 0
		for j := range mat[i] {
			cnt += mat[i][j]
		}
		arr[i] = cnt
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		a, b := arr[idx[i]], arr[idx[j]]
		return a < b || (a == b && idx[i] < idx[j])
	})

	return idx[:k]
}

// Use binary search + hashmap:
// m: matrix row
// n: matrix coloumn
// Time Complexity: O(m*logn + k)
// Space Complexity:O(m)
// Runtime: 12 ms, faster than 46.75%
// Memory Usage: 5.05 MB, less than 20.78%
func kWeakestRows2(mat [][]int, k int) []int {
	hashmap := map[int][]int{}
	result := []int{}
	calRow := func(i int) {
		cur := -1
		l, r := 0, len(mat[i])-1
		for l <= r {
			mid := l + (r-l)/2
			if mat[i][mid] == 0 {
				r = mid - 1
			} else {
				cur = mid
				l = mid + 1
			}
		}
		if v, ok := hashmap[cur+1]; ok {
			v = append(v, i)
			hashmap[cur+1] = v
		} else {
			hashmap[cur+1] = []int{i}
		}
	}

	for i := range mat {
		calRow(i)
	}

	cnt := 0
	for k > 0 {
		if v, ok := hashmap[cnt]; ok {
			if len(v) > k {
				v = v[:k]
			}
			result = append(result, v...)
			k -= len(v)

		}
		cnt++
	}

	return result
}
