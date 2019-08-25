/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.1
 */
public class ReverseLinkedList_206 {

    // Iterative version
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37.8 MB, less than 26.20%
    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }

        ListNode mNodeCurrent, mNodePrevious, mNodePrecede;

        mNodePrevious = null;
        mNodeCurrent = head;
        mNodePrecede = head.next;
        while (mNodePrecede != null) {
            mNodeCurrent.next = mNodePrevious;
            mNodePrevious = mNodeCurrent;
            mNodeCurrent = mNodePrecede;
            mNodePrecede = mNodePrecede.next;
        }
        mNodeCurrent.next = mNodePrevious;

        return mNodeCurrent;
    }

    // Recursive version
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37.9 MB, less than 13.75%
    public ListNode reverseList2(ListNode head) {
        if (head == null || head.next == null) { // empty list or 1 element list
            return head;
        }

        ListNode mNodePrecede = head.next;
        ListNode mNodeNew = reverseList2(mNodePrecede);

        mNodePrecede.next = head;
        head.next = null;
        return mNodeNew;
    }

    // Iterative version 2
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37 MB, less than 98.92%
    public ListNode reverseList3(ListNode head) {
        ListNode mCurrent = head, mPrevious = null;

        while (mCurrent != null) {
            ListNode mTmp = mCurrent.next;
            mCurrent.next = mPrevious;
            mPrevious = mCurrent;
            mCurrent = mTmp;
        }

        return mPrevious;
    }
}
