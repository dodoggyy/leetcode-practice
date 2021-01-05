package hard

import "math"

// Use BFS:
// n: wordList length
// C: word average length in wordList
// Time Complexity: O(n*C^2)
// Space Complexity:O(n*C)
// Runtime: 32 ms, faster than 90.23%
// Memory Usage: 9.2 MB, less than 14.29%
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordID := map[string]int{}
	graph := [][]int{}

	addWord := func(word string) int {
		id, exist := wordID[word]
		if !exist {
			id = len(wordID)
			wordID[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, v := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = v
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginID := addEdge(beginWord)
	endID, exist := wordID[endWord]
	if !exist {
		return 0
	}

	inf := math.MaxInt32
	dist := make([]int, len(wordID))
	for i := range dist {
		dist[i] = inf
	}
	dist[beginID] = 0
	queue := []int{beginID}
	for len(queue) > 0 {
		curID := queue[0]
		queue = queue[1:]
		if curID == endID {
			return dist[endID]/2 + 1
		}
		for _, v := range graph[curID] {
			if dist[v] == inf {
				dist[v] = dist[curID] + 1
				queue = append(queue, v)
			}
		}
	}
	return 0
}
