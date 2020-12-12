package medium

// Use Trie:
// K : length of the key
// Time Complexity: O(K)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.6 MB, less than 60.00%

type MapSum struct {
	hashmap map[string]int
	root    *TrieNode
}

type TrieNode struct {
	child map[rune]*TrieNode
	value int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		hashmap: make(map[string]int),
		root: &TrieNode{
			child: make(map[rune]*TrieNode),
			value: 0,
		},
	}
}

func (this *MapSum) Insert(key string, val int) {
	diff := val - this.hashmap[key]
	this.hashmap[key] = val
	cur := this.root

	for _, ch := range key {
		if _, exist := cur.child[ch]; !exist {
			cur.child[ch] = &TrieNode{child: make(map[rune]*TrieNode)}
		}
		cur = cur.child[ch]
		cur.value += diff
	}
}

func (this *MapSum) Sum(prefix string) int {
	cur := this.root
	for _, ch := range prefix {
		if _, exist := cur.child[ch]; !exist {
			return 0
		}
		cur = cur.child[ch]
	}
	return cur.value
}
