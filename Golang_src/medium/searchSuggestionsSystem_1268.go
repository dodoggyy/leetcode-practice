package medium

// Use trie:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 44 ms, faster than 50.00%
// Memory Usage: 8.91 MB, less than 92.35%
func suggestedProducts(products []string, searchWord string) [][]string {
	maxResult := 3
	result := [][]string{}

	trie := &Tries{}

	for _, v := range products {
		trie.Insert(v)
	}

	for i := range searchWord {
		result = append(result, trie.Search(searchWord[:i+1], maxResult))
	}

	return result
}

type Tries struct {
	isWord bool
	child  [26]*Tries
}

func ConstructorTries() Tries {
	return Tries{}
}

func (this *Tries) Insert(str string) {
	cur := this
	for _, ch := range str {
		tmp := ch - 'a'
		if cur.child[tmp] == nil {
			cur.child[tmp] = &Tries{}
		}
		cur = cur.child[tmp]
	}
	cur.isWord = true
}

func (this *Tries) Search(word string, maxResult int) []string {
	result := []string{}
	node := this.searchNode(word)
	if node == nil {
		return result
	}

	var dfs func(node *Tries, arr []byte)
	dfs = func(node *Tries, arr []byte) {
		if len(result) == maxResult {
			return
		}

		if node.isWord {
			result = append(result, string(arr))
		}

		for i, v := range node.child {
			if v != nil {
				dfs(v, append(arr, 'a'+byte(i)))
			}
		}
	}

	dfs(node, []byte(word))

	return result
}

func (this *Tries) searchNode(str string) *Tries {
	cur := this

	for _, ch := range str {
		tmp := ch - 'a'
		if cur.child[tmp] == nil {
			return nil
		}
		cur = cur.child[tmp]
	}

	return cur
}
