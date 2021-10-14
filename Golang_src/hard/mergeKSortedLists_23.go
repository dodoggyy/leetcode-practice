package hard

type ListNode struct {
	Val  int
	Next *ListNode
}

// Use divide and conquer:
// Time Complexity: O(knlogk)
// Space Complexity: O(logk)
// k: length of lists
// Runtime: 19 ms, faster than 45.41%
// Memory Usage: 6.3 MB, less than 36.39%
func mergeKLists2(lists []*ListNode) *ListNode {
	var mergeTwoList func(l1, l2 *ListNode) *ListNode
	mergeTwoList = func(l1, l2 *ListNode) *ListNode {
		if l1 == nil {
			return l2
		}
		if l2 == nil {
			return l1
		}
		if l1.Val < l2.Val {
			l1.Next = mergeTwoList(l1.Next, l2)
			return l1
		} else {
			l2.Next = mergeTwoList(l1, l2.Next)
			return l2
		}
	}

	var merge func(l, r int) *ListNode
	merge = func(l, r int) *ListNode {
		if l > r {
			return nil
		}
		if l == r {
			return lists[l]
		}
		mid := l + (r-l)/2
		return mergeTwoList(merge(l, mid), merge(mid+1, r))
	}

	return merge(0, len(lists)-1)
}
