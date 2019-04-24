/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class RemoveDuplicatesfromSortedList_83 {

    // Iterative version
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37.8 MB, less than 25.29%
    public ListNode deleteDuplicates(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }

        ListNode mNodeIndex = head;
        while (mNodeIndex.next != null) {
            if (mNodeIndex.val == mNodeIndex.next.val) {
                mNodeIndex.next = mNodeIndex.next.next;
            } else {
                mNodeIndex = mNodeIndex.next;
            }
        }

        return head;
    }

    // Recursive version
    // Runtime: 1 ms, faster than 30.22%
    // Memory Usage: 37.5 MB, less than 66.44%
    public ListNode deleteDuplicates2(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }
        head.next = deleteDuplicates2(head.next);
        return (head.val == head.next.val)? head.next: head;
    }
}
