package medium

// Use Search for a key prefix in a trie:
// Time Complexity: O(l)
// Space Complexity:O(n*l^2)
// Runtime: 52 ms, faster than 95.79%
// Memory Usage: 17.5 MB, less than 43.69%

type Trie struct {
	isWord bool
	child  [26]*Trie
}

/** Initialize your data structure here. */
func ConstructorTrie() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (root *Trie) Insert(word string) {
	cur := root
	for _, ch := range word {
		index := ch - 'a'
		if cur.child[index] == nil {
			cur.child[index] = &Trie{}
		}
		cur = cur.child[index]
	}
	cur.isWord = true
}

/** Returns if the word is in the trie. */
func (root *Trie) Search(word string) bool {
	cur := root.searchPrfix(word)
	return cur != nil && cur.isWord == true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (root *Trie) StartsWith(prefix string) bool {
	return root.searchPrfix(prefix) != nil
}

func (root *Trie) searchPrfix(word string) *Trie {
	cur := root
	for _, ch := range word {
		letter := ch - 'a'
		if cur.child[letter] == nil {
			return nil
		}
		cur = cur.child[letter]
	}

	return cur
}
