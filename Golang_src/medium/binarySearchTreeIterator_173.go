package medium

// Use Flattening the BST:
// Time Complexity: O(n) for initialization
// Time Complexity: O(1) for invoke
// Space Complexity:O(n)
// Runtime: 33 ms, faster than 57.09%
// Memory Usage: 11.4 MB, less than 5.67%
type BSTIterator struct {
	val []int
	idx int
}

func ConstructorBSTIterator(root *TreeNode) BSTIterator {
	var inorder func(*TreeNode) []int
	inorder = func(node *TreeNode) []int {
		if node == nil {
			return []int{}
		}

		tmp := []int{}
		tmp = append(tmp, inorder(node.Left)...)
		tmp = append(tmp, node.Val)
		tmp = append(tmp, inorder(node.Right)...)

		return tmp
	}

	return BSTIterator{val: inorder(root), idx: 0}
}

func (this *BSTIterator) Next() int {
	if this.idx < len(this.val) {
		this.idx++
		return this.val[this.idx-1]
	}
	return this.val[this.idx-1]
}

func (this *BSTIterator) HasNext() bool {
	return this.idx < len(this.val)
}

// Use iteration via stack:
// Time Complexity: O(1) for initialization
// Time Complexity: O(1) for invoke
// Space Complexity:O(n)
// Runtime: 35 ms, faster than 50.61%
// Memory Usage: 11.4 MB, less than 5.67%
type BSTIterator2 struct {
	stack []*TreeNode
	cur   *TreeNode
}

func ConstructorBSTIterator2(root *TreeNode) BSTIterator2 {
	return BSTIterator2{cur: root}
}

func (this *BSTIterator2) Next2() int {
	for node := this.cur; node != nil; node = node.Left {
		this.stack = append(this.stack, node)
	}
	this.cur = this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	val := this.cur.Val
	this.cur = this.cur.Right

	return val
}

func (this *BSTIterator2) HasNext2() bool {
	return this.cur != nil || len(this.stack) > 0
}
