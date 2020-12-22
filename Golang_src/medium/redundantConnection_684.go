package medium

// Use Union-Find:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 4 ms, faster than 96.77%
// Memory Usage: 3.1 MB, less than 100.00%
type UnionFindSet struct {
	parents []int
	size    int
}

func New(size int) *UnionFindSet {
	parent := make([]int, size+1)
	for index := range parent {
		parent[index] = index
	}
	return &UnionFindSet{parents: parent, size: size + 1}
}

func (u *UnionFindSet) Root(p int) int {
	for p != u.parents[p] {
		p = u.parents[p]
	}
	return p
}

func (u *UnionFindSet) Connected(p, q int) bool {
	return u.Root(p) == u.Root(q)
}

func (u *UnionFindSet) Union(p, q int) {
	rootP := u.Root(p)
	rootQ := u.Root(q)
	u.parents[rootQ] = rootP
}

func findRedundantConnection(edges [][]int) []int {
	result := []int{}
	u := New(len(edges))
	for _, edge := range edges {
		if u.Connected(edge[0], edge[1]) {
			result = edge
		} else {
			u.Union(edge[0], edge[1])
		}
	}

	return result
}
