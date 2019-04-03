/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class LinkedListCycle_141 {

    class ListNode {
        int val;
        ListNode next;

        ListNode(int x) {
            val = x;
            next = null;
        }
    }

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
}
