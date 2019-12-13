/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class OddEvenLinkedList_328 {
    // 2019年12月14日
    // Use pointers
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 36.3 MB, less than 100.00%
    public ListNode oddEvenList(ListNode head) {
        if (head == null) {
            return head;
        }
        ListNode mOdd = head, mEven = head.next, mEvenHead = mEven;
        while (mEven != null && mEven.next != null) {
            mOdd.next = mOdd.next.next;
            mEven.next = mEven.next.next;
            mOdd = mOdd.next;
            mEven = mEven.next;
        }
        mOdd.next = mEvenHead;
        return head;
    }
}
