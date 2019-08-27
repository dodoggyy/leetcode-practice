/**
 * 
 */
package com.medium;

import java.util.HashSet;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class LinkedListCycleII_142 {
    // 2019年8月27日
    // Use HashSet
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 6 ms, faster than 19.79%
    // Memory Usage: 36.4 MB, less than 6.32%
    public ListNode detectCycle(ListNode head) {
        HashSet<ListNode> mSet = new HashSet<>();
        while (head != null) {
            if (mSet.contains(head)) {
                return head;
            } else {
                mSet.add(head);
                head = head.next;
            }
        }
        return null;
    }

    // 2019年8月27日
    // Use two pointers
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 34.5 MB, less than 95.79%
    public ListNode detectCycle2(ListNode head) {
        // L1: distance from start to cycle entry
        // L2: distance from cycle entry to meeting point
        // C: cycle length
        // mSlow: L1 + L2
        // mFast: (L1 + L2) * 2
        // => L1 + L2 + N*C = (L1 + L2) *2
        // => L1 = C - L2
        ListNode mFast = head;
        ListNode mSlow = head;

        do {
            if (mFast == null || mFast.next == null) {
                return null;
            }

            mFast = mFast.next.next;
            mSlow = mSlow.next;
        } while (mFast != mSlow);
        mFast = head;
        while (mFast != mSlow) {
            mFast = mFast.next;
            mSlow = mSlow.next;
        }

        return mFast;

    }
}
