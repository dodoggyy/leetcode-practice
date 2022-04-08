package easy

import (
	"container/heap"
	"sort"
)

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

func (h *hp) push(v int) {
	heap.Push(h, v)
}

func (h *hp) pop() int {
	return heap.Pop(h).(int)
}

// Use max heap:
// Time Complexity: O(nlogn)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.9 MB, less than 82.69%
func lastStoneWeight(stones []int) int {
	q := &hp{stones}
	heap.Init(q)
	for q.Len() > 1 {
		x, y := q.pop(), q.pop()
		q.push(x - y)
	}

	if q.Len() > 0 {
		return q.IntSlice[0]
	}

	return 0
}
