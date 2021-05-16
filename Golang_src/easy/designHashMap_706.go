package easy

import "container/list"

// Use double linkedlist:
// Time Complexity: O(n/b)
// Space Complexity:O(n+b)
// n: total elemets of hash set
// b: total chain
// Runtime: 116 ms, faster than 57.58%
// Memory Usage: 8 MB, less than 49.24%
const baseHash = 769

type entry struct {
	key, val int
}

type MyHashMap struct {
	data []list.List
}

/** Initialize your data structure here. */
func ConstructorHashMap() MyHashMap {
	return MyHashMap{
		data: make([]list.List, baseHash),
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % baseHash
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			e.Value = entry{
				key: key,
				val: value,
			}
			return
		}
	}
	this.data[h].PushBack(entry{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			return e.Value.(entry).val
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			this.data[h].Remove(e)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
