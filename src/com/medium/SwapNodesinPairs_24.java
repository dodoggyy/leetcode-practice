/**
 * 
 */
package com.medium;

import com.easy.ListNode;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SwapNodesinPairs_24 {

    // Iterative version
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 35.8 MB, less than 91.13%
    public ListNode swapPairs(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }
        ListNode mNodeDummy = new ListNode(0);
        mNodeDummy.next = head;
        ListNode mNodePrevious = mNodeDummy;

        while (mNodePrevious.next != null && mNodePrevious.next.next != null) {
          ListNode mNodeTmp = mNodePrevious.next.next;
          mNodePrevious.next.next = mNodeTmp.next;
          mNodeTmp.next = mNodePrevious.next;
          mNodePrevious.next = mNodeTmp;
          mNodePrevious = mNodeTmp.next;
        }

        return mNodeDummy.next;
    }
    
    // Recursive version
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 35.7 MB, less than 91.13% 
    public ListNode swapPairs2(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }
        // O1(head) -> O2 -> O3 -> null
        // swap O2 and O3
        ListNode mNodeTmp = head.next; // tmp node point to 2nd node
        head.next = swapPairs2(head.next.next); // 1st node point to 3rd
        mNodeTmp.next = head; // 3rd node point to 2nd
        
        return mNodeTmp;
    }
}
