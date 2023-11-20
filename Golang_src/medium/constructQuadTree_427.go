package medium

// Definition for a QuadTree node.
type QuadNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadNode
	TopRight    *QuadNode
	BottomLeft  *QuadNode
	BottomRight *QuadNode
}

// Use Optimized Recursion:
// Time Complexity: O(n^2)
// Space Complexity:O(n^2)
// Runtime: 4 ms, faster than 96.00%
// Memory Usage: 13.84 MB, less than 40.00%
func construct(grid [][]int) *QuadNode {
	n := len(grid)
	pre := make([][]int, n+1)
	pre[0] = make([]int, n+1)
	for i := range grid {
		pre[i+1] = make([]int, n+1)
		for j, v := range grid[i] {
			pre[i+1][j+1] = pre[i+1][j] + pre[i][j+1] - pre[i][j] + v
		}
	}

	var dfs func(r0, c0, r1, c1 int) *QuadNode
	dfs = func(r0, c0, r1, c1 int) *QuadNode {
		total := pre[r1][c1] - pre[r1][c0] - pre[r0][c1] + pre[r0][c0]
		if total == 0 { // leaf
			return &QuadNode{Val: false, IsLeaf: true}
		}
		if total == (r1-r0)*(c1-c0) {
			return &QuadNode{Val: true, IsLeaf: true}
		}
		midR, midC := (r0+r1)/2, (c0+c1)/2
		return &QuadNode{
			Val:         true,
			IsLeaf:      false,
			TopLeft:     dfs(r0, c0, midR, midC),
			TopRight:    dfs(r0, midC, midR, c1),
			BottomLeft:  dfs(midR, c0, r1, midC),
			BottomRight: dfs(midR, midC, r1, c1),
		}
	}

	return dfs(0, 0, n, n)
}
