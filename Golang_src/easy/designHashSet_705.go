package easy

import "container/list"

// Use list:
// Time Complexity: O(n/b)
// Space Complexity:O(n+b)
// n: total elemets of hash set
// b: total chain
// Runtime: 76 ms, faster than 93.48%
// Memory Usage: 14.2 MB, less than 6.52%
const base = 769

type MyHashSet struct {
	data []list.List
}

/** Initialize your data structure here. */
func ConstructorHashSet() MyHashSet {
	return MyHashSet{
		data: make([]list.List, base),
	}
}

func (this *MyHashSet) hash(key int) int {
	return key % base
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		h := this.hash(key)
		this.data[h].PushBack(key)
	}
}

func (this *MyHashSet) Remove(key int) {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			this.data[h].Remove(e)
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
