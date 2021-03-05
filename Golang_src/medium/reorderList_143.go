package medium

// Use linear array store:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 8 ms, faster than 96.75%
// Memory Usage: 6 MB, less than 33.77%
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	arr := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		arr = append(arr, node)
	}
	indexStart, indexEnd := 0, len(arr)-1

	for indexStart < indexEnd {
		arr[indexStart].Next = arr[indexEnd]
		indexStart++
		if indexStart == indexEnd {
			break
		}
		arr[indexEnd].Next = arr[indexStart]
		indexEnd--
	}
	arr[indexStart].Next = nil
}

// find mid, half inverse, merge list:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 96.75%
// Memory Usage: 5.4 MB, less than 42.21%
func reorderList2(head *ListNode) {
	if head == nil {
		return
	}
	mid := findMidNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)

	mergeList(l1, l2)
}

func findMidNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}

func mergeList(l1, l2 *ListNode) {
	var next1, next2 *ListNode
	for l1 != nil && l2 != nil {
		next1 = l1.Next
		next2 = l2.Next

		l1.Next = l2
		l1 = next1

		l2.Next = l1
		l2 = next2
	}
}
