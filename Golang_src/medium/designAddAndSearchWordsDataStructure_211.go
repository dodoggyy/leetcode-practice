package medium

// Use trie:
// AddWord:
// m for length of word
// Time Complexity: O(m)
// Space Complexity:O(m)
// Search:
// Time Complexity: O(m)
// Space Complexity:O(1)
// Runtime: 68 ms, faster than 86.11%
// Memory Usage: 20 MB, less than 20.37%
type WordDictionary struct {
	isWord bool
	child  [26]*WordDictionary
}

/** Initialize your data structure here. */
func ConstructorWordDictionary() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this
	for _, ch := range word {
		tmp := ch - 'a'
		if cur.child[tmp] == nil {
			cur.child[tmp] = &WordDictionary{}
		}
		cur = cur.child[tmp]
	}
	cur.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	cur := this
	for idx, ch := range word {
		if ch == '.' {
			for _, dict := range cur.child {
				if dict != nil {
					if dict.Search(word[idx+1:]) {
						return true
					}
				}
			}
			return false
		} else {
			tmp := ch - 'a'
			if cur.child[tmp] == nil {
				return false
			}
			cur = cur.child[tmp]
		}

	}

	return cur.isWord == true
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
