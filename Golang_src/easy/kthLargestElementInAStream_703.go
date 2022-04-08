package easy

import (
	"container/heap"
	"sort"
)

// Use priority queue:
// Time Complexity: O(nlogk)
// Space Complexity:O(1ogk)
// Runtime: 57 ms, faster than 42.06%
// Memory Usage: 7.8 MB, less than 69.84%
type KthLargest struct {
	sort.IntSlice
	k int
}

func ConstructorPriorityQueue(k int, nums []int) KthLargest {
	pq := KthLargest{k: k}
	for _, v := range nums {
		pq.Add(v)
	}

	return pq
}

func (this *KthLargest) Push(val interface{}) {
	this.IntSlice = append(this.IntSlice, val.(int))
}

func (this *KthLargest) Pop() interface{} {
	v := this.IntSlice[len(this.IntSlice)-1]
	this.IntSlice = this.IntSlice[:len(this.IntSlice)-1]
	return v
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.k {
		heap.Pop(this)
	}
	return this.IntSlice[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
