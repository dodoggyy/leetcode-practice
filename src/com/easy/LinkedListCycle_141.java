/**
 * 
 */
package com.easy;

import java.util.HashSet;

/**
 * @author Chris Lin
 * @version 1.1
 */
public class LinkedListCycle_141 {

    // Use two pointer
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 38.2 MB, less than 93.78%
    public boolean hasCycle(ListNode head) {
        if (head == null) {
            return false;
        }
        // use 2 node, one for 1 step and the other for 2 step
        ListNode mNode1 = head;
        ListNode mNode2 = head.next;

        // if cycle exist, 2 pointer may equal
        while ((mNode1 != null) && (mNode2 != null) && (mNode2.next != null)) {
            if (mNode1 == mNode2) {
                return true;
            }
            mNode1 = mNode1.next;
            mNode2 = mNode2.next.next;
        }

        return false;
    }

    // 2019年8月27日
    // Use two pointer 2
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 36.9 MB, less than 100.00%
    public boolean hasCycle2(ListNode head) {
        if (head == null) {
            return false;
        }
        ListNode mFast = head.next;
        ListNode mSlow = head;

        while (mFast != mSlow) {
            if (mFast == null || mFast.next == null) {
                return false;
            }
            mFast = mFast.next.next;
            mSlow = mSlow.next;
        }
        return true;
    }

    // 2019年8月27日
    // Use HashSet
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 4 ms, faster than 20.20%
    // Memory Usage: 37.1 MB, less than 100.00%
    public boolean hasCycle3(ListNode head) {
        HashSet<ListNode> mSet = new HashSet<>();

        while (head != null) {
            if (mSet.contains(head)) {
                return true;
            } else {
                mSet.add(head);
                head = head.next;
            }
        }

        return false;
    }
}
