package medium

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Use recursive:
// Time Complexity: O(nlogn)
// Space Complexity:O(logn)
// Runtime: 548 ms, faster than 8.47%
// Memory Usage: 191.4 MB, less than 83.90%
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	mid := findMid(head)

	node := &TreeNode{Val: mid.Val}

	if head == mid {
		return node
	}

	node.Left = sortedListToBST(head)
	node.Right = sortedListToBST(mid.Next)

	return node
}

func findMid(head *ListNode) *ListNode {
	slow := head
	fast := head
	var pre *ListNode = nil

	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// handle slow = head
	if pre != nil {
		pre.Next = nil
	}

	return slow
}

// Use array & recursive:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 1028 ms, faster than 5.08%
// Memory Usage: 71.6 MB, less than 100.00%
func sortedListToBST2(head *ListNode) *TreeNode {
	num := []int{}

	for head != nil {
		num = append(num, head.Val)
		head = head.Next
	}

	return toBST(num, 0, len(num)-1)
}

func toBST(num []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	node := &TreeNode{Val: num[mid]}
	node.Left = toBST(num, start, mid-1)
	node.Right = toBST(num, mid+1, end)

	return node
}
